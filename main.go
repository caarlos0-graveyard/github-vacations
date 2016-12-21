package main

import (
	"os"

	"fmt"

	"strings"

	"github.com/caarlos0/spin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	var token = os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Missing GITHUB_TOKEN environment variable")
		os.Exit(2)
	}
	if len(os.Args) != 2 {
		fmt.Println("Missing organization name to ignore")
		os.Exit(2)
	}
	var ignoring = strings.ToLower(os.Args[1])

	spin := spin.New("%s Helping you to not work...")
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
	}
	var count int
	for _, notification := range notifications {
		owner := *notification.Repository.Owner.Login
		if strings.ToLower(owner) == ignoring {
			client.Activity.DeleteThreadSubscription(*notification.ID)
			client.Activity.MarkThreadRead(*notification.ID)
			count++
		}
	}
	spin.Stop()
	fmt.Println(count, "notifications marked as read")
}
