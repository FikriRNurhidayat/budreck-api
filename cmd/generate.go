package cmd

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	MIGRATION_KIND = "migrations"
)

func Generate(kind string, name string) {
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	switch kind {
	case MIGRATION_KIND:
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), name)
		os.Create(fmt.Sprintf("%s/db/migrations/%s.up.sql", workdir, filename))
		os.Create(fmt.Sprintf("%s/db/migrations/%s.down.sql", workdir, filename))
	}
}
