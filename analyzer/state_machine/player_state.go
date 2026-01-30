package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) player(playerId models.PlayerId) *PlayerState {
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

	OnSnail            bool
	LastRecordedSnailX int
	BeingEaten         bool

	HasSpeed  bool
	IsWarrior bool

	IsBot   bool
	IsQueen bool
}

func (s *PlayerState) respawn() {
	s.HasBerry = false
	s.OnSnail = false
	s.BeingEaten = false

	s.HasSpeed = false
	s.IsWarrior = false

	// Don't touch IsBot or IsQueen
}

// stateArrayIdx converts speed/drone/warrior/queen state into a unique int
// vanilla drone is 0
// speed drone is 1
// vanilla warrior is 2
// speed warrior is 3
// queen is 4
func (s *PlayerState) stateArrayIdx() int {
	if s.IsQueen {
		return 4
	} else if s.IsWarrior && s.HasSpeed {
		return 3
	} else if s.IsWarrior {
		return 2
	} else if s.HasSpeed {
		return 1
	} else {
		return 0
	}
}
