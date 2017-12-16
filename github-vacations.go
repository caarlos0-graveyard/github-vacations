// Package githubvacations contains functions that help you not work
// when you are on vacations (or just don't want to).
package githubvacations

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Notification from github
type Notification struct {
	URL, Title, Repo, Reason string
}

// MarkWorkNotificationsAsRead checks your notifications from work and mark
// them as read
func MarkWorkNotificationsAsRead(token, org string) ([]Notification, error) {
	var result []Notification
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
		return result, err
	}
	for _, notification := range notifications {
		var owner = notification.GetRepository().GetOwner().GetLogin()
		if strings.ToLower(owner) == strings.ToLower(org) {
			if _, err = client.Activity.DeleteThreadSubscription(ctx, notification.GetID()); err != nil {
				return result, err
			}
			// if _, err = client.Activity.MarkThreadRead(ctx, notification.GetID()); err != nil {
			// 	return result, err
			// }
			var url = notification.GetSubject().GetURL()
			for old, new := range map[string]string{
				"api.github.com": "github.com",
				"/repos/":        "/",
				"/pulls/":        "/pull/",
			} {
				url = strings.Replace(url, old, new, 1)
			}
			result = append(result, Notification{
				URL:    url,
				Title:  notification.GetSubject().GetTitle(),
				Repo:   notification.GetRepository().GetFullName(),
				Reason: notification.GetReason(),
			})
		}
	}
	return result, err
}
