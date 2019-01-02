package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/softleader/homebrew-tap/tapper/pkg/brew"
	"github.com/softleader/homebrew-tap/tapper/pkg/gh"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type tapperServerCmd struct {
	repo *gh.Repo
	port string
}

func main() {
	var verbose bool
	c := tapperServerCmd{
		repo: &gh.Repo{},
	}
	cmd := &cobra.Command{
		Use:   "tapper-server",
		Short: "tapper-server is a bot automatic upgrade Homebrew Formula",
		Long:  "tapper-server is a bot automatic upgrade Homebrew Formula",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			c.repo.Owner = os.ExpandEnv(c.repo.Owner)
			c.repo.Name = os.ExpandEnv(c.repo.Name)
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
	f.StringVar(&c.repo.Owner, "owner", "$GITHUB_OWNER", "github owner. Overrides $GITHUB_OWNER")
	f.StringVar(&c.repo.Name, "repo", "$GITHUB_REPO", "github repo. Overrides $GITHUB_REPO")
	f.StringVar(&c.port, "port", "$PORT", "server port")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (cmd *tapperServerCmd) run() error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		sh, err := ioutil.ReadFile("./script/get.sh")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.String(http.StatusOK, string(sh))
		}
	})
	r.GET("/tapper", func(c *gin.Context) {
		c.File("tapper.tgz")
	})
	r.PUT("/:formula/:version", func(c *gin.Context) {
		formula := &brew.Formula{
			Name:    c.Param("formula"),
			Version: c.Param("version"),
		}
		token := c.GetHeader("GITHUB_TOKEN")
		if err := formula.Upgrade(token, cmd.repo); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.Status(http.StatusOK)
		}
	})
	r.Run(fmt.Sprintf(":%s", cmd.port))

	return nil
}
