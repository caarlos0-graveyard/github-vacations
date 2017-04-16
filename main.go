// Package githubvacations contains functions that help you not work
// when you are on vacations (or just don't want to).
package githubvacations

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// MarkWorkNotificationsAsRead checks your notifications from work and mark
// them as read
func MarkWorkNotificationsAsRead(token, org string) (count int, err error) {
	var ctx = context.Background()
	var client = github.NewClient(oauth2.NewClient(
		ctx,
		oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})),
	)
	notifications, _, err := client.Activity.ListNotifications(
		ctx,
		&github.NotificationListOptions{},
	)
	if err != nil {
		return
	}
	for _, notification := range notifications {
		var owner = *notification.Repository.Owner.Login
		if strings.ToLower(owner) == org {
			client.Activity.DeleteThreadSubscription(ctx, *notification.ID)
			client.Activity.MarkThreadRead(ctx, *notification.ID)
			count++
		}
	}
	return
}
