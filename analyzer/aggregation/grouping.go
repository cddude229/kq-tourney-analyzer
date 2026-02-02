package aggregation

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
)

type StateMachineGrouping struct {
	StateMachine *models.StateMachine
	TourneyMatch hivemind.TourneyMatch
	GameId       int64
}
