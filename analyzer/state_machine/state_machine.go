package state_machine

import (
	"cddude229/kq-tourney-analyzer/models"
)

type StateMachine struct {
	playerState map[models.PlayerId]*PlayerState
	playerStats map[models.PlayerId]*PlayerStats

	blueBerries      int
	goldBerries      int
	remainingBerries int

	winningTeam  *models.TeamColor2
	winCondition *models.WinCondition
}

func New() *StateMachine {
	return &StateMachine{
		playerState: make(map[models.PlayerId]*PlayerState),
		playerStats: make(map[models.PlayerId]*PlayerStats),
	}
}
