package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "scanbot",
	Short: "CLI tool for interacting with and processing ImunifyAV malware scans",
	Long:  "This tool allows you to interact with ImunifyAV using the command line. ImunifyAV has command line tools, but these can be clunky to use. scanbot provides a simple list of commands for the most common tasks, and allows you to present and export the results in machine-readable formats.",
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
