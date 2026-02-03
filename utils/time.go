package utils

import (
	"fmt"
	"strings"
	"time"
)

func ConvertToTimestamp(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	layouts := []string{
		"02/01/2006 15:04:05",
		"02/01/2006 15:04",
		"02/01/2006",
	}
	var t time.Time
	var err error
	for _, l := range layouts {
		t, err = time.Parse(l, dateStr) // time.Parse -> trả UTC
		if err == nil {
			return t.UTC(), nil
		}
	}
	return time.Time{}, fmt.Errorf("cannot parse time %q", dateStr)
}
