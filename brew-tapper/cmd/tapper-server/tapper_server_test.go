package main

import (
	"github.com/softleader/homebrew-tap/tapper/pkg/brew"
	"github.com/softleader/homebrew-tap/tapper/pkg/gh"
	"os"
	"testing"
)

func TestUpgradeFormula(t *testing.T) {
	var found bool
	repo := &gh.Repo{}
	var token string
	if token, found = os.LookupEnv("GITHUB_TOKEN"); !found {
		t.SkipNow()
	}
	if repo.Name, found = os.LookupEnv("GITHUB_REPO"); !found {
		t.SkipNow()
	}
	if repo.Owner, found = os.LookupEnv("GITHUB_OWNER"); !found {
		t.SkipNow()
	}
	formula := &brew.Formula{
		Name:    "slctl",
		Version: "0.1.9",
	}
	if err := formula.Upgrade(token, repo); err != nil {
		t.Error(err)
	}
}
