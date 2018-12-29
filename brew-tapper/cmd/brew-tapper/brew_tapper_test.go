package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestUpgradeFormula(t *testing.T) {
	var found bool
	gh := &gh{}
	var token string
	if token, found = os.LookupEnv("GITHUB_TOKEN"); !found {
		t.SkipNow()
	}
	if gh.repo, found = os.LookupEnv("GITHUB_REPO"); !found {
		t.SkipNow()
	}
	if gh.owner, found = os.LookupEnv("GITHUB_OWNER"); !found {
		t.SkipNow()
	}
	formula := &formula{
		name:    "slctl",
		version: "0.1.9",
	}
	if err := upgradeFormula(token, gh, formula); err != nil {
		t.Error(err)
	}
}

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
	formula := formula{
		version: "9.9.9",
	}
	out := format(lua, &formula)
	if !strings.Contains(out, fmt.Sprintf("version %q", formula.version)) {
		t.Errorf("version should be %s", formula.version)
	}
	fmt.Println(out)
}
