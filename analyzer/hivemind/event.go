package hivemind

import (
	"time"
)

type HivemindEvent struct {
	Id        int64
	Timestamp time.Time
	EventType string
	Values    string
	GameId    int64
}

type ById []HivemindEvent

// TODO: Sorting UTs
func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }
