package database

import (
	"scanbot/util"
)

type Scan struct {
	ScanId         string
	Started        string
	Completed      string
	Path           string
	TotalMalicious int
}

type MalwareHit struct {
	File string
	Type string
	Size string
	Date string
}

const getScansQuery string = `
SELECT scanid, started, completed, path, total_malicious
FROM malware_scans
WHERE completed NOT NULL
ORDER BY started DESC
LIMIT ?`

// GetScans queries the database for a specified number of scans, sorted latest first.
func GetScans(count int) ([]Scan, error) {
	rows, err := database.Query(getScansQuery, count)
	if err != nil {
		return nil, err
	}
	var scans []Scan
	for rows.Next() {
		var scan Scan
		err := rows.Scan(&scan.ScanId, &scan.Started, &scan.Completed, &scan.Path, &scan.TotalMalicious)
		if err != nil {
			return nil, err
		}
		scans = append(scans, scan)
	}
	return scans, nil
}

const getScanQuery string = `
SELECT scanid, started, completed, path, total_malicious
FROM malware_scans
WHERE scanid = ?`

// GetScan queries the database for a specific scan.
func GetScan(scanId string) (Scan, error) {
	row := database.QueryRow(getScanQuery, scanId)
	var scan Scan
	err := row.Scan(&scan.ScanId, &scan.Started, &scan.Completed, &scan.Path, &scan.TotalMalicious)
	if err != nil {
		return scan, err
	}
	return scan, nil
}

const getMalwareHitsQuery string = `
SELECT orig_file, type, size, CAST(timestamp AS INT)
FROM malware_hits
WHERE scanid_id = ?
ORDER BY timestamp DESC
`

// GetMalwareHits queries the database for the malware hits of a specific scan.
func GetMalwareHits(scanId string) ([]MalwareHit, error) {
	rows, err := database.Query(getMalwareHitsQuery, scanId)
	if err != nil {
		return nil, err
	}
	var hits []MalwareHit
	for rows.Next() {
		var hit MalwareHit
		err := rows.Scan(&hit.File, &hit.Type, &hit.Size, &hit.Date)
		if err != nil {
			return nil, err
		}
		hit.Date = util.FormatTimestamp(hit.Date)
		hits = append(hits, hit)
	}
	return hits, nil
}
