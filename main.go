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
	"sync"
)

var currentVersion = "v0.0.0"

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

	var wg sync.WaitGroup
	wg.Add(1)

	var updateStatus updater.UpdateStatus
	var updateErr error

	go func() {
		updateStatus, updateErr = u.Update()
		wg.Done()
	}()

	wg.Wait()

	if updateErr != nil {
		color.Red("Scanbot was unable to update. Be aware there may be bugs in this version, check https://github.com/BenTomsett/scanbot for updates.")
	}

	if updateStatus == updater.Updated {
		color.Green("Scanbot has been updated to the latest version and will now close. Re-run your command to continue.")
		os.Exit(1)
	}
}

func main() {
	c := color.New(color.BgGreen)
	c.Printf(" SCANBOT %s \n\n", currentVersion)

	selfUpdate()

	database.Connect()
	cmd.Execute()
}
