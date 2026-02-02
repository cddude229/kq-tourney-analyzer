package models

func (s *StateMachine) Stats(playerId PlayerId) *PlayerStats {
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

	// State tracking
	SpeedDroneUptime     int64
	VanillaWarriorUptime int64
	SpeedWarriorUptime   int64
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

func sumCounterByNestedKey(counter [][]int, nestedKeys ...playerClass) int {
	total := 0
	for _, row := range counter {
		for _, nestedKey := range nestedKeys {
			total += row[nestedKey]
		}
	}
	return total
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
	return sumCounterByNestedKey(s.KillCounter, classQueen, classSpeedWarrior, classWarrior, classSpeedDrone, classVanillaDrone)
}

func (s *PlayerStats) TotalDeaths() int {
	return sumCounterByNestedKey(s.DeathCounter, classQueen, classSpeedWarrior, classWarrior, classSpeedDrone, classVanillaDrone)
}

func (s *PlayerStats) MilKills() int {
	return sumCounterByNestedKey(s.KillCounter, classQueen, classSpeedWarrior, classWarrior)
}

func (s *PlayerStats) MilDeaths() int {
	return sumCounterByNestedKey(s.DeathCounter, classQueen, classSpeedWarrior, classWarrior)
}

func (s *PlayerStats) QueenKills() int {
	return sumCounterByNestedKey(s.KillCounter, classQueen)
}
