package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
	"os"
	"runtime"
	"scanbot/cmd"
	"scanbot/database"
)

var currentVersion = "0.0.1"

func selfUpdate() {
	p := &provider.Github{
		RepositoryURL: "https://github.com/BenTomsett/scanbot",
	}
	latestVersion, err := p.GetLatestVersion()
	if err != nil {
		color.Red("Unable to check for updates. Be aware there may be bugs in this version, check https://github.com/BenTomsett/scanbot for updates.")
	}
	p.ArchiveName = fmt.Sprintf("scanbot-%s-%s-%s.tar.gz", latestVersion, runtime.GOOS, runtime.GOARCH)

	u := &updater.Updater{
		Provider:       p,
		ExecutableName: "scanbot",
		Version:        currentVersion,
	}

	if updateAvailable, err := u.CanUpdate(); err != nil {
		color.Red("Unable to check for updates. Be aware there may be bugs in this version, check https://github.com/BenTomsett/scanbot for updates.")
	} else if updateAvailable {
		status, err := u.Update()

		if err != nil {
			color.Red("Scanbot was unable to automatically update. Be aware there may be bugs in this version, check https://github.com/BenTomsett/scanbot for updates.")
		}

		if status == updater.Updated {
			color.Green("Scanbot has updated to the latest version and will now exit. Re-run your command.")
			os.Exit(0)
		}
	}
}

func main() {
	selfUpdate()

	c := color.New(color.BgGreen)
	c.Printf(" SCANBOT v%s \n\n", currentVersion)
	database.Connect()
	cmd.Execute()
}
