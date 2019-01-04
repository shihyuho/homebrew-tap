package brew

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"github.com/softleader/homebrew-tap/tapper/pkg/gh"
	"golang.org/x/oauth2"
	"regexp"
	"time"
)

const (
	rb = "Formula/%s.rb"
)

var (
	author             = "softleader/homebrew-tap/brew-tapper"
	mail               = "supprt@softleader.com.tw"
	msg                = "version upgrade by brew-tapper bot"
	versionRegexp      = regexp.MustCompile(`(version\s)(.+)`)
	sha256Regexp       = regexp.MustCompile(`(sha256\s)(.+)`)
	darwinSha256Regexp = regexp.MustCompile(`(OS\.mac[\s|\S]+sha256\s)[\S|\s]+elsif`)
	linuxSha256Regexp  = regexp.MustCompile(`(OS\.linux[\s|\S]+sha256\s)[\S|\s]+end`)
)

func (f *Formula) Upgrade(token string, repo *gh.Repo) error {
	logrus.Printf("upgrading %s formula %q to %q", repo, f.Name, f.Version)
	ctx := context.Background()
	client := newTokenClient(ctx, token)
	filePath := fmt.Sprintf(rb, f.Name)
	fileContent, _, _, err := client.Repositories.GetContents(ctx, repo.Owner, repo.Name, filePath, &github.RepositoryContentGetOptions{})
	if err != nil {
		return err
	}
	content, err := fileContent.GetContent()
	if err != nil {
		return err
	}
	upgraded := format(content, f)

	now := time.Now()
	author := &github.CommitAuthor{
		Name:  &author,
		Email: &mail,
		Date:  &now,
	}
	opt := &github.RepositoryContentFileOptions{}
	opt.Message = &msg
	opt.Content = []byte(upgraded)
	opt.SHA = fileContent.SHA
	opt.Author = author
	opt.Committer = author
	_, _, err = client.Repositories.UpdateFile(ctx, repo.Owner, repo.Name, filePath, opt)
	return err
}

func newTokenClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func format(origin string, f *Formula) (out string) {
	logrus.Debugf("formatting formula:\n%s", origin)
	out = versionRegexp.ReplaceAllString(origin, fmt.Sprintf("$1%q", f.Version))
	out = darwinSha256Regexp.ReplaceAllStringFunc(out, func(s string) string {
		return sha256Regexp.ReplaceAllString(s, fmt.Sprintf("$1%q", f.DarwinSha256))
	})
	out = linuxSha256Regexp.ReplaceAllStringFunc(out, func(s string) string {
		return sha256Regexp.ReplaceAllString(s, fmt.Sprintf("$1%q", f.LinuxSha256))
	})
	logrus.Debugf("version replaced:\n%s", out)
	return
}
