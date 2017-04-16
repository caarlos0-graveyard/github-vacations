package main

import (
	"fmt"
	"os"
	"strings"

	githubvacations "github.com/caarlos0/github-vacations"
	"github.com/caarlos0/spin"
	"github.com/urfave/cli"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "github-vacations"
	app.Usage = "Automagically ignore all notifications related to work when you are on vacations"
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "org, o",
			Usage: "Organization name to ignore",
		},
		cli.StringFlag{
			Name:   "token, t",
			EnvVar: "GITHUB_TOKEN",
			Usage:  "GitHub token",
		},
	}
	app.Action = func(c *cli.Context) error {
		var token = c.String("token")
		var org = strings.ToLower(c.String("org"))
		if token == "" {
			return cli.NewExitError("missing GITHUB_TOKEN", 1)
		}
		if org == "" {
			return cli.NewExitError("missing organization to ignore", 1)
		}
		var spin = spin.New("%s Helping you to not work...")
		spin.Start()
		count, err := githubvacations.MarkWorkNotificationsAsRead(token, org)
		spin.Stop()
		if err != nil {
			return cli.NewExitError("failed to mark notifications as read", 1)
		}
		fmt.Printf("%v notifications marked as read", count)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
