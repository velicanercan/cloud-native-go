package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/velicanercan/cloud-native-go/config"
)

const (
	dialect = "pgx"
)

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String("dir", "migrations", "directory with migration files")
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])
	args := flags.Args()
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}

	command := args[0]

	c := config.NewDB()
	dbString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.Username, c.Password, c.DBName)
	fmt.Println(dbString)
	db, err := goose.OpenDBWithDriver(dialect, dbString)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	if err := goose.Run(command, db, *dir, args[1:]...); err != nil {
		log.Fatalf("migrate %v:%v", command, err)
	}
}

func usage() {
	log.Print(usagePrefix)
	flags.PrintDefaults()
	log.Print(usageCommands)
}

var (
	usagePrefix = `Usage: migrate COMMAND
Examples:
	migrate status
	migrate up
	migrate down
`
	usageCommands = `
Commands:
	up                   Migrate the DB to the most recent version available
	up-by-one            Migrate the DB up by 1
	up-to VERSION        Migrate the DB to a specific VERSION
	down                 Roll back the version by 1
	down-to VERSION      Roll back to a specific VERSION
	redo                 Re-run the latest migration
	reset				 Roll back all migrations
	status               Dump the migration status for the current DB
	version              Print the current version of the database
	create NAME [sql|go] Creates new migration file with the current timestamp
	fix 				 Apply sequential ordering to migrations
`
)
