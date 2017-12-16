package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/boltdb/bolt"
	lib "github.com/caarlos0/github-vacations"
)

func init() {
	log.SetHandler(cli.Default)
}

const (
	version = "master"
)

var (
	app    = kingpin.New("github-vacations", "Automagically ignore all notifications related to work when you are on vacations")
	dbPath = app.Flag("db", "Path to the database").Default("$HOME/.vacations.db").String()
	check  = app.Command("check", "Check for work notifications and mark them as read").Default()
	org    = check.Flag("org", "Organization name to ignore").Required().Short('o').String()
	token  = check.Flag("token", "Your GitHub token").Required().Short('t').Envar("GITHUB_TOKEN").String()
	read   = app.Command("read", "Read notifications that were previously marked as read")
)

func main() {
	fmt.Printf("\n")
	defer fmt.Printf("\n")
	app.Version(version)
	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case check.FullCommand():
		checkNotifications(*dbPath, *org, *token)
	case read.FullCommand():
		readNotification(*dbPath)
	}
}

func readNotification(path string) {
	db, err := bolt.Open(os.ExpandEnv(path), 0600, nil)
	kingpin.FatalIfError(err, "%v")
	if err := db.View(func(tx *bolt.Tx) error {
		var bucket = tx.Bucket([]byte("notifications"))
		if bucket == nil {
			log.Info("0 notifications to read")
			return nil
		}

		bucket.ForEach(func(url, title []byte) error {
			printLink(title, url)
			return nil
		})
		return nil
	}); err != nil {
		kingpin.FatalIfError(err, "%v")
	}
}

func printLink(title, url []byte) {
	// TODO: this is a hack because printf won't accept `\e`
	var e = string(0x1b)
	log.Infof("%s]8;;%s\a%s%s]8;;\a", e, url, title, e)
}

func checkNotifications(path, org, token string) {
	log.Info("helping you not to work...")
	db, err := bolt.Open(os.ExpandEnv(path), 0600, nil)
	kingpin.FatalIfError(err, "%v")
	defer db.Close()

	notifications, err := lib.MarkWorkNotificationsAsRead(token, org)
	kingpin.FatalIfError(err, "%v")
	log.Infof("%d %s notifications mark as read", len(notifications), org)

	if err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("notifications"))
		if err != nil {
			return err
		}
		for _, notification := range notifications {
			bucket.Put([]byte(notification.URL), []byte(notification.Title))
		}
		return nil
	}); err != nil {
		kingpin.FatalIfError(err, "%v")
	}
	log.Infof("notifications stored on %s", path)
}
