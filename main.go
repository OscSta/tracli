package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
)

func main() {

	db, err := sql.Open("sqlite3", "./testdb.db")
	if err != nil {
		log.Fatal("Error at open: ", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Error at ping: ", err)
	}

	app := &cli.App{
		Name:    "tracli",
		Version: "0.1",
		Usage:   "Track stuff",
		Authors: []*cli.Author{
			{
				Name:  "Oscar Stålnacke",
				Email: "oscstal@gmail.com",
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		return
	}
}
