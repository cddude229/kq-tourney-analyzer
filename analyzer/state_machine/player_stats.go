package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) stats(playerId models.PlayerId) *PlayerStats {
	stats, ok := s.PlayerStats[playerId]
	if !ok {
		stats = &PlayerStats{
			BumpCounter:  makeEmptyCounter(),
			KillCounter:  makeEmptyCounter(),
			DeathCounter: makeEmptyCounter(),
		}
		s.PlayerStats[playerId] = stats
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

func (s *PlayerStats) TotalBerries() int {
	return s.BerriesDunked + s.BerriesKickedOurTeam
}

func (s *PlayerStats) TotalKills() int {
	totalKills := 0
	for _, row1 := range s.KillCounter {
		for _, row2 := range row1 {
			totalKills += row2
		}
	}
	return totalKills
}

func (s *PlayerStats) TotalDeaths() int {
	totalDeaths := 0
	for _, row1 := range s.DeathCounter {
		for _, row2 := range row1 {
			totalDeaths += row2
		}
	}
	return totalDeaths
}

func (s *PlayerStats) MilKills() int {
	milKills := 0
	for _, row1 := range s.KillCounter {
		milKills += row1[4] // Queen
		milKills += row1[3] // SW
		milKills += row1[2] // VW
	}
	return milKills
}

func (s *PlayerStats) MilDeaths() int {
	milDeaths := 0
	for _, row1 := range s.DeathCounter {
		milDeaths += row1[4] // Queen
		milDeaths += row1[3] // SW
		milDeaths += row1[2] // VW
	}
	return milDeaths
}

func (s *PlayerStats) QueenKills() int {
	milKills := 0
	for _, row1 := range s.KillCounter {
		milKills += row1[4]
	}
	return milKills

}
