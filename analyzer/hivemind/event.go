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

type ByTimestamp []HivemindEvent

// TODO: Sorting UTs
func (a ByTimestamp) Len() int           { return len(a) }
func (a ByTimestamp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTimestamp) Less(i, j int) bool { return a[i].Timestamp.Compare(a[j].Timestamp) == -1 }
