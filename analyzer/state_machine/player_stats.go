package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) stats(playerId models.PlayerId) *PlayerStats {
	stats, ok := s.playerStats[playerId]
	if !ok {
		stats = &PlayerStats{}
		s.playerStats[playerId] = stats
	}
	return stats
}

type PlayerStats struct {
	// Berries
	BerriesDunked          int
	BerriesKickedOurTeam   int
	BerriesKickedTheirTeam int

	// Gate usage
	GateDenyKills int
	KilledInGate  int
	LeftGate      int
}
