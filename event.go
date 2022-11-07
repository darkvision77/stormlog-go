package stormlog

import (
	"fmt"
	"time"
)

const RFC3339Nano = "2006-01-02T15:04:05.000000000Z07:00"

type Event struct {
//	Id        int
	Level     Level
	Module    string
//	Thread    int
	Timestamp time.Time
	Message   string
}

func (e Event) String() string {
	return fmt.Sprintf("%s [%s] %s: %s", e.Timestamp.Format(time.RFC3339Nano), e.Level.String(), e.Module, e.Message)
}
