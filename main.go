package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/spin"
	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
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

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(oauth2.NoContext, ts)
		client := github.NewClient(tc)
		notifications, _, err := client.Activity.ListNotifications(
			&github.NotificationListOptions{},
		)
		if err != nil {
			spin.Stop()
			return cli.NewExitError(err.Error(), 1)
		}
		var count int
		for _, notification := range notifications {
			owner := *notification.Repository.Owner.Login
			if strings.ToLower(owner) == org {
				client.Activity.DeleteThreadSubscription(*notification.ID)
				client.Activity.MarkThreadRead(*notification.ID)
				count++
			}
		}
		spin.Stop()
		fmt.Println(count, "notifications marked as read")
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
