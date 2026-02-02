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

	HasSpeed  bool
	IsWarrior bool

	IsBot   bool
	IsQueen bool
}

func (s *PlayerState) respawn() {
	s.HasBerry = false
	s.OnSnail = false
	s.IsEating = false
	s.BeingEaten = false

	s.HasSpeed = false
	s.IsWarrior = false

	// Don't touch IsBot or IsQueen
}

// stateArrayIdx converts speed/drone/warrior/queen state into a unique int
func (s *PlayerState) stateArrayIdx() playerClass {
	if s.IsQueen {
		return classQueen
	} else if s.IsWarrior && s.HasSpeed {
		return classSpeedWarrior
	} else if s.IsWarrior {
		return classWarrior
	} else if s.HasSpeed {
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
