package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	rb = "Formula/%s.rb"
)

var (
	author        = "softleader/homebrew-tap/tapper"
	mail          = "supprt@softleader.com.tw"
	msg           = fmt.Sprintf("version upgrade by %s", author)
	versionRegexp = regexp.MustCompile("(version )(.+)")
)

type fisherCmd struct {
	gh   *gh
	port string
}

type gh struct {
	token, owner, repo string
}

type formula struct {
	name, version string
}

type sha struct {
	darwin, linux, windows string
}

func (gh *gh) newTokenClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gh.token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func main() {
	var verbose bool
	c := fisherCmd{
		gh: &gh{},
	}
	cmd := &cobra.Command{
		Use:   "fisher",
		Short: "Fisher is a bot automatic update Fish-Food",
		Long:  "Fisher is a bot automatic update Fish-Food",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			c.gh.token = os.ExpandEnv(c.gh.token)
			c.gh.owner = os.ExpandEnv(c.gh.owner)
			c.gh.repo = os.ExpandEnv(c.gh.repo)
			c.port = os.ExpandEnv(c.port)
			if _, err := strconv.Atoi(c.port); err != nil {
				c.port = strconv.Itoa(8080)
			}
			logrus.SetOutput(cmd.OutOrStdout())
			if verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
			return c.run()
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&verbose, "verbose", "v", verbose, "enable verbose output")
	f.StringVar(&c.gh.token, "token", "$GITHUB_TOKEN", "github access token. Overrides $GITHUB_TOKEN")
	f.StringVar(&c.gh.owner, "owner", "$GITHUB_OWNER", "github owner. Overrides $GITHUB_OWNER")
	f.StringVar(&c.gh.repo, "repo", "$GITHUB_REPO", "github repo. Overrides $GITHUB_REPO")
	f.StringVar(&c.port, "port", "$PORT", "server port")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (cmd *fisherCmd) run() error {
	r := gin.Default()
	r.PUT("/:formula", func(c *gin.Context) {
		formula := &formula{
			name:    c.Param("formula"),
			version: c.Request.FormValue("version"),
		}
		if err := upgradeFormula(cmd.gh, formula); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.Status(http.StatusOK)
		}
	})
	r.Run(fmt.Sprintf(":%s", cmd.port))

	return nil
}

func upgradeFormula(gh *gh, formula *formula) error {
	ctx := context.Background()
	client := gh.newTokenClient(ctx)
	foodPath := fmt.Sprintf(rb, formula.name)
	fileContent, _, _, err := client.Repositories.GetContents(ctx, gh.owner, gh.repo, foodPath, &github.RepositoryContentGetOptions{})
	if err != nil {
		return err
	}
	content, err := fileContent.GetContent()
	if err != nil {
		return err
	}
	upgraded := format(content, formula)

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
	_, _, err = client.Repositories.UpdateFile(ctx, gh.owner, gh.repo, foodPath, opt)
	return err
}

func format(origin string, food *formula) (out string) {
	out = versionRegexp.ReplaceAllString(origin, fmt.Sprintf("$1%q", food.version))
	return
}
