package state_machine

import (
	"cddude229/kq-tourney-analyzer/models"
	"time"
)

type GateStateAndStats struct {
	Color      *models.TeamColor1
	LastTagged time.Time

	TimeUntagged int64
	TimeForGold  int64
	TimeForBlue  int64
}
