package models

import (
	"time"
)

type GateStateAndStats struct {
	Color      *TeamColor1
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
