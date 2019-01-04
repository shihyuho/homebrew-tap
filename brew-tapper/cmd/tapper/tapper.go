package main

import (
	"github.com/sirupsen/logrus"
	"github.com/softleader/homebrew-tap/tapper/pkg/brew"
	"github.com/softleader/homebrew-tap/tapper/pkg/gh"
	"github.com/spf13/cobra"
	"os"
)

type tapperCmd struct {
	repo    *gh.Repo
	token   string
	dist    string
	formula *brew.Formula
}

func main() {
	var verbose bool
	c := tapperCmd{
		repo:    &gh.Repo{},
		formula: &brew.Formula{},
	}
	cmd := &cobra.Command{
		Use:   "tapper",
		Short: "tapper is a bot to upgrade Homebrew Formula",
		Long:  "tapper is a bot to upgrade Homebrew Formula",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logrus.SetOutput(cmd.OutOrStdout())
			if verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
			return c.run()
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&verbose, "verbose", "v", verbose, "enable verbose output")
	f.StringVar(&c.repo.Owner, "owner", "", "github owner")
	f.StringVar(&c.repo.Name, "repo", "homebrew-tap", "github repo")
	f.StringVar(&c.token, "token", "", "github token")
	f.StringVar(&c.dist, "dist", "_dist", "dist folder to store binary archive")
	f.StringVar(&c.formula.Name, "formula", "", "homebrew formula")
	f.StringVar(&c.formula.Version, "version", "", "homebrew formula version to upgrade to")

	cmd.MarkFlagRequired("owner")
	cmd.MarkFlagRequired("token")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (cmd *tapperCmd) run() error {
	if cmd.formula.NotSpecified() {
		if err := cmd.formula.Guess(cmd.dist); err != nil {
			return err
		}
	}
	return cmd.formula.Upgrade(cmd.token, cmd.repo)
}
