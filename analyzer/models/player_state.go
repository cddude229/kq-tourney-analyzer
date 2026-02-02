package models

import (
	"time"
)

func (s *StateMachine) player(playerId PlayerId) *PlayerState {
	playerState, ok := s.playerState[playerId]
	if !ok {
		playerState = &PlayerState{
			IsQueen: playerId.IsQueen(),
		}
		playerState.respawn()
		s.playerState[playerId] = playerState
	}
	return playerState
}

type PlayerState struct {
	HasBerry bool

	OnSnail               bool
	LastRecordedSnailX    int
	LastRecordedSnailTime time.Time
	IsEating              bool
	BeingEaten            bool

	IsSpeed         bool
	GotSpeedAt      time.Time
	IsWarrior       bool
	BecameWarriorAt time.Time

	IsBot   bool
	IsQueen bool
}

func (s *PlayerState) respawn() {
	s.HasBerry = false
	s.OnSnail = false
	s.IsEating = false
	s.BeingEaten = false

	s.IsSpeed = false
	s.IsWarrior = false

	// Don't touch IsBot or IsQueen
}

// stateArrayIdx converts speed/drone/warrior/queen state into a unique int
func (s *PlayerState) stateArrayIdx() playerClass {
	if s.IsQueen {
		return classQueen
	} else if s.IsWarrior && s.IsSpeed {
		return classSpeedWarrior
	} else if s.IsWarrior {
		return classWarrior
	} else if s.IsSpeed {
		return classSpeedDrone
	} else {
		return classVanillaDrone
	}
}

type playerClass int

const (
	classVanillaDrone playerClass = iota
	classSpeedDrone
	classWarrior
	classSpeedWarrior
	classQueen
)
