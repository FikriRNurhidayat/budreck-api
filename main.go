package main

import (
	"errors"
	"log"
	"os"

	"budreck/cmd"
)

func main() {
	if err := validate(os.Args); err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}

	run(os.Args)
}

func validate(args []string) (err error) {
	if len(args) < 2 {
		err = errors.New("sub-command is required")
	}

	return
}

func run(args []string) {
	switch args[1] {
	case "serve":
		cmd.Serve()
	case "db:migrate":
		cmd.Migrate(cmd.MIGRATE_UP)
	case "db:rollback":
		cmd.Migrate(cmd.MIGRATE_DOWN)
	case "generate":
		handleGenerateCMD(args)
	default:
		os.Exit(0)
	}
}

func handleGenerateCMD(args []string) {
	switch args[2] {
	case "migration":
		cmd.Generate(cmd.MIGRATION_KIND, args[3])
	}
}
