package hivemind

import (
	"cddude229/kq-tourney-analyzer/models"
	"time"
)

type TourneyMatch struct {
	Id                int64
	StartTime         time.Time
	EndTime           time.Time
	WinCondition      models.WinCondition
	WinningTeam       string
	MapName           string
	PlayerCount       int64
	CabinetId         int64
	CabinetName       string
	TournamentMatchId int64
	BlueTeamName      string
	GoldTeamName      string
}

type TourneyMatchById []TourneyMatch

// TODO: Sorting UTs
func (a TourneyMatchById) Len() int           { return len(a) }
func (a TourneyMatchById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TourneyMatchById) Less(i, j int) bool { return a[i].StartTime.Before(a[j].StartTime) }
