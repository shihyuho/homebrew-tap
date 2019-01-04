package brew

import (
	"fmt"
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	lua := `class Slctl < Formula
  desc "Slctl is a command line interface for running commands against SoftLeader Services"
  homepage "https://github.com/softleader/slctl"
  version "2.1.0"
  
  if OS.mac?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-darwin-#{version}.tgz"
	sha256 "-----"
  elsif OS.linux?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-linux-#{version}.tgz"
	sha256 "-----"
  end

  depends_on :arch => :x86_64
  
  def install
    bin.install "slctl"
  end

  def caveats; <<~EOS
    To begin working with slctl, run the 'slctl init' command.
  EOS
  end
end`
	formula := Formula{
		Version:      "9.9.9",
		DarwinSha256: "abc",
		LinuxSha256:  "def",
	}
	out := format(lua, &formula)
	if !strings.Contains(out, fmt.Sprintf("version %q", formula.Version)) {
		t.Errorf("version should be %s", formula.Version)
	}
	if !strings.Contains(out, fmt.Sprintf("sha256 %q", formula.DarwinSha256)) {
		t.Errorf("darwin sha256 should be %q", formula.DarwinSha256)
	}
	if !strings.Contains(out, fmt.Sprintf("sha256 %q", formula.LinuxSha256)) {
		t.Errorf("linux sha256 should be %q", formula.LinuxSha256)
	}
	fmt.Println(out)
}
