package brew

import (
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var (
	extensions = []string{
		".zip",
		".tar",
		".tar.gz",
		".tgz",
		".tar.bz2",
		".tbz2",
		".tar.xz",
		".txz",
		".tar.lz4",
		".tlz4",
		".tar.sz",
		".tsz",
		".rar",
		".bz2",
		".gz",
		".lz4",
		".sz",
		".xz",
	}
	goarch64bit = []string{
		"amd64",
		"arm64",
		"arm64be",
		"ppc64",
		"ppc64le",
		"mips64",
		"mips64le",
		"s390x",
		"sparc64",
	}
	goos = []string{
		"android",
		"darwin",
		"dragonfly",
		"freebsd",
		"linux",
		"nacl",
		"netbsd",
		"openbsd",
		"plan9",
		"solaris",
		"windows",
		"zos",
	}
)

func (f *Formula) Guess(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() || !isSupportedArchive(file.Name()) {
			continue
		}
		bestGuess, err := guess(filepath.Base(file.Name()))
		if err != nil {
			logrus.Debugln(err)
			continue
		}
		if len(f.Name) == 0 {
			f.Name = bestGuess.Name
		}
		if len(f.Version) == 0 {
			f.Version = bestGuess.Version
		}
		return nil
	}

	return fmt.Errorf("not found any binary archive")
}

func isSupportedArchive(source string) bool {
	for _, suffix := range extensions {
		if strings.HasSuffix(source, suffix) {
			return true
		}
	}
	return false
}

func guess(file string) (*Formula, error) {
	var name = truncateExtension(file)
	chunks := strings.Split(name, "-")
	if len(chunks) == 1 {
		if chunks = strings.Split(file, "_"); len(chunks) == 1 {
			return nil, fmt.Errorf("%s does not contains any dash or underline, consider it not a binary archive", file)
		}
	}
	f := &Formula{}

	for _, chunk := range chunks {
		if _, err := semver.NewVersion(chunk); err == nil {
			f.Version = chunk
			continue
		}
		if containsIgnoreCase(goarch64bit, chunk) {
			continue
		}
		if containsIgnoreCase(goos, chunk) {
			continue
		}
		f.Name = chunk
	}
	return f, nil
}

func containsIgnoreCase(ss []string, s string) bool {
	s = strings.ToLower(s)
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

func truncateExtension(file string) string {
	var ext = filepath.Ext(file)
	return file[0 : len(file)-len(ext)]
}
