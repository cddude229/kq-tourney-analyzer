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

func (s *StateMachine) Gate(x int, y int) *GateStateAndStats {
	key := x * y
	gate, exists := s.gates[key]
	if !exists {
		gate = &GateStateAndStats{}
		s.gates[key] = gate
	}
	return gate
}
