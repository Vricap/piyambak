package utils

import (
	"fmt"
	"time"
)

var WIB, _ = time.LoadLocation("Asia/Jakarta")

func FormatTimestamp(timestamp string) (string, error) {
	layouts := []string{
		time.RFC3339,          // 2026-08-31T09:29:08Z
		"2006-01-02 15:04:05", // 2026-08-31 09:29:08
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, timestamp)
		if err == nil {
			// Convert UTC -> WIB (UTC+7)
			t = t.In(WIB)
			return fmt.Sprintf("%s WIB", t.Format("2 January. 15:04")), nil
		}
	}

	return "", fmt.Errorf("invalid timestamp format: %s", timestamp)
}
