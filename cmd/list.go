package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"scanbot/database"
	"scanbot/util"
)

// RenderScanTable takes a list of scans, and renders them in a table. The parameter showIndex determines whether a
// column is shown numbering the rows.
func RenderScanTable(scans []database.Scan, showIndex bool) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	var tbl table.Table
	if showIndex {
		tbl = table.New("#", "Scan ID", "Started", "Completed", "Path", "# Malicious files")
	} else {
		tbl = table.New("Scan ID", "Started", "Completed", "Path", "# Malicious files")
	}

	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	var row = 1

	for _, scan := range scans {
		if showIndex {
			tbl.AddRow(row, scan.ScanId, util.FormatTimestamp(scan.Started), util.FormatTimestamp(scan.Completed), scan.Path, scan.TotalMalicious)
		} else {
			tbl.AddRow(scan.ScanId, util.FormatTimestamp(scan.Started), util.FormatTimestamp(scan.Completed), scan.Path, scan.TotalMalicious)
		}
		row++
	}

	tbl.Print()
}

// handleList is the main function for the list command
func handleList(cmd *cobra.Command, args []string) {
	count, _ := cmd.Flags().GetInt("count")
	fmt.Printf("Getting the %d most recent scans.\n\n", count)
	scans, err := database.GetScans(count)
	if err != nil {
		color.Red("An error occurred while retrieving the latest scans.")
		panic(err)
	}
	RenderScanTable(scans, false)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Return a list of the latest malware scans, with the most recent first.",
	Run:   handleList,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntP("count", "c", 10, "The number of results to return.")
}
