package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("Hi")
	app := &cli.App{Name: "tracli", Usage: "Track stuff"}
	app.Run(os.Args)
}
