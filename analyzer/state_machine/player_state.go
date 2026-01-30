package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) player(playerId models.PlayerId) *PlayerState {
	playerState, ok := s.playerState[playerId]
	if !ok {
		playerState = &PlayerState{}
		playerState.respawn()
		s.playerState[playerId] = playerState
	}
	return playerState
}

type PlayerState struct {
	HasBerry   bool
	OnSnail    bool
	BeingEaten bool

	HasSpeed  bool
	IsWarrior bool

	IsBot bool
}

func (s *PlayerState) respawn() {
	s.HasBerry = false
	s.OnSnail = false
	s.BeingEaten = false

	s.HasSpeed = false
	s.IsWarrior = false

	// Don't touch IsBot
}
