package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) stats(playerId models.PlayerId) *PlayerStats {
	stats, ok := s.playerStats[playerId]
	if !ok {
		stats = &PlayerStats{
			BumpCounter:  makeEmptyCounter(),
			KillCounter:  makeEmptyCounter(),
			DeathCounter: makeEmptyCounter(),
		}
		s.playerStats[playerId] = stats
	}
	return stats
}

type PlayerStats struct {
	// Berries
	BerriesDunked          int
	BerriesKickedOurTeam   int
	BerriesKickedTheirTeam int

	// Snail
	SnailDistance int

	// Gate usage
	GateDenyKills int
	KilledInGate  int
	LeftGate      int

	// Bumps - maps our state vs their state when we bump
	BumpCounter [][]int

	// In both cases, maps killer's state to victim's state
	KillCounter  [][]int
	DeathCounter [][]int
}

func (s *PlayerStats) TotalKills() int {
	return 0 // TODO
}

func makeEmptyCounter() [][]int {
	return [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
}

func (s *PlayerStats) recordSnailDistance(dist int) {
	// TODO: This is technically not correct since the snail can be bumped backward.
	if dist < 0 {
		s.SnailDistance -= dist
	} else {
		s.SnailDistance += dist
	}
}
