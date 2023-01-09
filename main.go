package main

import (
	"github.com/fatih/color"
	"scanbot/cmd"
	"scanbot/database"
)

var version = "0.0.1"

func main() {
	c := color.New(color.BgGreen)
	c.Printf(" SCANBOT v%s \n\n", version)
	database.Connect()
	cmd.Execute()
}
