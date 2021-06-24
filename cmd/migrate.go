package cmd

import (
	"fmt"
	"log"
	"os"

	"budreck/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	MIGRATE_DOWN string = "DOWN"
	MIGRATE_UP   string = "UP"
)

func Migrate(direction string) (err error) {
	var path string
	var m *migrate.Migrate

	if path, err = os.Getwd(); err != nil {
		log.Fatal(err.Error())
		return
	}

	if m, err = migrate.New(
		fmt.Sprintf("file://%s/db/migrations", path),
		config.GetDatabaseConnectionURL()); err != nil {
		log.Fatal(err.Error())
		return
	}

	switch direction {
	case "UP":
		if err = m.Up(); err != nil {
			log.Fatal(err.Error())
			return
		}
	case "DOWN":
		if err = m.Steps(-1); err != nil {
			log.Fatal(err.Error())
			return
		}
	}

	return
}
