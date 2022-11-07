package stormlog_test

import (
	"testing"
	"time"

	"github.com/darkvision77/stormlog-go"
)

// 2022-11-08T01:30:17.29569572+03:00 [INFO] main: 1
// 2022-11-08T01:30:17.295740734+03:00 [INFO] main: 2

func comp(t *testing.T, timestamp time.Time) {
	const expectedFmt = "2006-01-02T15:04:05.000000000Z"
	s := timestamp.Format(stormlog.RFC3339Nano)
	t.Log(s)
	if len(s) != len(expectedFmt) {
		t.Errorf("Format mismatch: actual=%s", s)
	}
}

func TestTimeFormat(t *testing.T) {
	timestamp := time.Date(2022, 11, 8, 1, 30, 17, 70993435, time.UTC)
	comp(t, timestamp)
	comp(t, timestamp.Add(5 * time.Nanosecond))
}
