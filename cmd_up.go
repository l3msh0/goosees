package main

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"log"
)

var upCmd = &Command{
	Name:    "up",
	Usage:   "",
	Summary: "Migrate the DB to the most recent version available",
	Help:    `up extended help here...`,
	Run:     upRun,
}

func upRun(cmd *Command, args ...string) {

	confs, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	for _, conf := range confs {
		target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
		if err != nil {
			log.Fatal(err)
		}

		if err := goose.RunMigrations(conf, conf.MigrationsDir, target); err != nil {
			log.Fatal(err)
		}
	}
}
