package state_machine

import (
	"cddude229/kq-tourney-analyzer/models"
)

type StateMachine struct {
	playerState map[models.PlayerId]*PlayerState
	playerStats map[models.PlayerId]*PlayerStats

	gates map[int]*GateStateAndStats

	mapName     string
	goldOnLeft  bool
	attractMode bool
	cabVersion  *string

	blueBerries      int
	goldBerries      int
	remainingBerries int

	winningTeam       *models.TeamColor2
	winCondition      *models.WinCondition
	finalGameDuration float64 // uninitialized if not final value
}

func New() *StateMachine {
	return &StateMachine{
		playerState: make(map[models.PlayerId]*PlayerState),
		playerStats: make(map[models.PlayerId]*PlayerStats),

		gates: make(map[int]*GateStateAndStats),
	}
}
