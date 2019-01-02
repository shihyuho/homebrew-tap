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
  version "0.1.8"
  url "https://github.com/softleader/homebrew-tap/releases/download/slctl/slctl-darwin-#{version}.tgz"
  
  depends_on "git"

  def install
    bin.install "slctl"
  end

  def caveats; <<~EOS
    To begin working with slctl, run the 'slctl init' command.
  EOS
  end
end`
	formula := Formula{
		Version: "9.9.9",
	}
	out := format(lua, &formula)
	if !strings.Contains(out, fmt.Sprintf("version %q", formula.Version)) {
		t.Errorf("version should be %s", formula.Version)
	}
	fmt.Println(out)
}
