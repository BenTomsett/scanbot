package cmd

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"scanbot/database"
	"scanbot/util"
)

// createReport retrieves malware hits for a given scanId and saves to a CSV file
func createReport(scanId string, completed string) {
	color.Yellow("Retrieving malware hits for scan %s completed on %s.\n", scanId, util.FormatTimestamp(completed))

	hits, getHitsErr := database.GetMalwareHits(scanId)
	if getHitsErr != nil {
		color.Red("An error occurred while retrieving malware hits.")
		panic(getHitsErr)
	}

	count := len(hits)

	if count == 0 {
		color.Red("No malware hits found for the specified scan. ImunifyAV sometimes clears old hits from the database, so you may need to run a new scan if there are meant to be results.")
		os.Exit(1)
	}

	color.Yellow("Found %d malware hits.\n\n", count)

	reportFile, createReportErr := os.Create("imunify-report-" + completed + ".csv")
	if createReportErr != nil {
		color.Red("An error occurred while creating the report file.")
		panic(createReportErr)
	}

	writer := csv.NewWriter(reportFile)

	writeHeaderErr := writer.Write([]string{"path", "malwareType", "size", "datetime"})
	if writeHeaderErr != nil {
		color.Red("An error occurred while writing to the report file.")
		panic(writeHeaderErr)
	}

	for _, hit := range hits {
		writeRowErr := writer.Write([]string{hit.File, hit.Type, hit.Size, hit.Date})
		println(hit.File, hit.Type, hit.Size, hit.Date)
		if writeRowErr != nil {
			color.Red("An error occurred while writing to the report file.")
			panic(writeHeaderErr)
		}
	}

	writer.Flush()

	color.Green("Imunify scan report written to imunify-report-%s.csv", completed)

	closeFileErr := reportFile.Close()
	if closeFileErr != nil {
		color.Red("An error occurred while closing the report file.")
		panic(closeFileErr)
	}
}

// handleReport is the main function for the report command
func handleReport(cmd *cobra.Command, args []string) {
	var scanId string

	if len(args) == 0 {
		// No scanId provided, so get a list of scans and prompt the user to select one

		color.Yellow("No scan ID provided. Retrieving the most recent scans...\n")
		scans, err := database.GetScans(10)
		if err != nil {
			color.Red("An error occurred while retrieving the latest scans.")
			panic(err)
		}

		RenderScanTable(scans, true)

		var n int
		for {
			fmt.Print("Choose a scan from the above list: ")
			_, err := fmt.Scanf("%d\n", &n)
			if err != nil {
				color.Red("Enter the number shown next to the scan you want to generate a report for. If you don't see the scan you want, use the 'list' command to see all scans, and then use 'scanbot report <scanId>'.")
				fmt.Scanln()
				continue
			}
			if n < 1 || n > len(scans) {
				color.Red("Enter the number shown next to the scan you want to generate a report for. If you don't see the scan you want, use the 'list' command to see all scans, and then use 'scanbot report <scanId>'.")
				continue
			}
			n--
			println()
			break
		}
		scanId = scans[n].ScanId
	} else {
		// scanId provided in the args, so use that
		scanId = args[0]
	}

	// Get the scan details to make sure the specified scan exists
	scan, err := database.GetScan(scanId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			color.Red("No scan found with the specified ID.")
			os.Exit(1)
		default:
			color.Red("An error occurred while retrieving details of the specified scan.")
			panic(err)
		}
	}

	createReport(scan.ScanId, scan.Completed)
}

var reportCmd = &cobra.Command{
	Use:   "report <scan ID>",
	Short: "Creates a CSV report for a given scan ID.",
	Long:  `If provided with a scan ID, this command will generate a CSV report for the specified scan. If no scan ID is provided, the user will be prompted to choose a scan from the latest 10 scans.`,
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
	Run:   handleReport,
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
