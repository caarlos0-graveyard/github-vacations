package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/boltdb/bolt"
	lib "github.com/caarlos0/github-vacations"
	"github.com/fatih/color"
)

func init() {
	log.SetHandler(cli.Default)
}

const (
	version = "master"
	// TODO: this is a hack because printf won't accept `\e`
	backslashE = string(0x1b)
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
	if err != nil {
		log.WithError(err).Fatal("failed to open database")
	}
	defer closeDB(db)
	if err := db.View(func(tx *bolt.Tx) error {
		var bucket = tx.Bucket([]byte("notifications"))
		if bucket == nil {
			log.Warn("0 notifications to read")
			return nil
		}
		var notifications = map[string][]lib.Notification{}
		var count int
		if err := bucket.ForEach(func(url, encoded []byte) error {
			var notification lib.Notification
			if err := json.Unmarshal(encoded, &notification); err != nil {
				return err
			}
			count++
			notifications[notification.Repo] = append(notifications[notification.Repo], notification)
			return nil
		}); err != nil {
			log.WithError(err).Fatal("failed to read bucket")
		}
		var bold = color.New(color.Bold)
		log.Infof(bold.Sprintf("you have %d notifications in %d repositories!", count, len(notifications)))
		for repo, repoNotifications := range notifications {
			fmt.Printf("\n")
			cli.Default.Padding = 3
			log.Infof("%s:", repo)
			cli.Default.Padding = 6
			for _, n := range repoNotifications {
				log.Infof(
					"%s]8;;%s\a%s%s]8;;\a",
					backslashE, n.URL, n.Title, backslashE,
				)
			}
			cli.Default.Padding = 3
		}
		return nil
	}); err != nil {
		log.WithError(err).Fatal("failed to read database")
	}
}

func checkNotifications(path, org, token string) {
	log.Info("helping you not to work...")
	db, err := bolt.Open(os.ExpandEnv(path), 0600, nil)
	if err != nil {
		log.WithError(err).Fatal("failed to open database")
	}
	defer closeDB(db)

	notifications, err := lib.MarkWorkNotificationsAsRead(token, org)
	if err != nil {
		log.WithError(err).Fatal("failed check your notifications")
	}
	log.Infof("%d %s notifications mark as read", len(notifications), org)

	if err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("notifications"))
		if err != nil {
			return err
		}
		for _, notification := range notifications {
			encoded, err := json.Marshal(notification)
			if err != nil {
				return err
			}
			if err := bucket.Put([]byte(notification.URL), encoded); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		log.WithError(err).Fatal("failed to store your notifications")
	}
	log.Infof("notifications stored on %s", path)
}

func closeDB(db *bolt.DB) {
	if err := db.Close(); err != nil {
		log.WithError(err).Error("failed to close db")
	}
}
