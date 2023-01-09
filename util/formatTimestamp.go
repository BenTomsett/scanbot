package util

import (
	"strconv"
	"time"
)

// FormatTimestamp transforms a UNIX timestamp into human-readable format (RFC822, e.g. "01 Jan 23 15:00 GMT")
func FormatTimestamp(timestamp string) string {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return ""
	}
	t := time.Unix(i, 0)
	return t.Format(time.RFC822)
}
