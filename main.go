package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
)

const databasePath = "./traclidb.db"
const testTableName = "test"

type NHPage struct {
	Id int
}

func assertTableExists(db *sql.DB, tableName string) {
	stmt := fmt.Sprintf(
		`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER NOT NULL PRIMARY KEY
		);
		`,
		tableName,
	)
	_, err := db.Exec(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
	}
}

func addPageFromID(ctx *cli.Context) error {
	id, err := strconv.Atoi(ctx.Args().First())
	if err != nil {
		log.Fatal(err, " - Could not parse as integer")
	}
	// Open connection to database and defer closing
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal("Error at open: ", err)
	}
	defer db.Close()
	// Test database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error at ping: ", err)
	}
	assertTableExists(db, testTableName)
	insertStatement := fmt.Sprintf(`INSERT INTO %s values (%d)`, testTableName, id)
	_, err = db.Exec(insertStatement)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func main() {

	// Setup CLI app
	app := &cli.App{
		Name:    "tracli",
		Version: "0.1",
		Usage:   "Track stuff through the command line",
		Authors: []*cli.Author{
			{
				Name:  "Oscar Stålnacke",
				Email: "oscstal@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "test",
				Subcommands: []*cli.Command{
					{
						Name:   "add",
						Action: addPageFromID,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
