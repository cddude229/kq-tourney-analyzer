package models

func (s *StateMachine) Stats(playerId PlayerId) *PlayerStats {
	stats, ok := s.playerStats[playerId]
	if !ok {
		stats = &PlayerStats{
			BumpCounter:  makeEmptyCounter(),
			KillCounter:  makeEmptyCounter(),
			DeathCounter: makeEmptyCounter(),

			GamesPlayed: 1,
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

	// State tracking
	SpeedDroneUptime     int64
	VanillaWarriorUptime int64
	SpeedWarriorUptime   int64

	// Bumps - maps our state vs their state when we bump
	BumpCounter [][]int

	// In both cases, maps killer's state to victim's state
	KillCounter  [][]int
	DeathCounter [][]int

	GamesPlayed int
}

func (p *PlayerStats) Merge(other ...*PlayerStats) *PlayerStats {
	newStats := &PlayerStats{
		BumpCounter:  makeEmptyCounter(),
		KillCounter:  makeEmptyCounter(),
		DeathCounter: makeEmptyCounter(),
	}

	other = append(other, p)
	for _, oldStats := range other {
		newStats.BerriesDunked += oldStats.BerriesDunked
		newStats.BerriesKickedOurTeam += oldStats.BerriesKickedOurTeam
		newStats.BerriesKickedTheirTeam += oldStats.BerriesKickedTheirTeam

		newStats.SnailDistance += oldStats.SnailDistance

		newStats.GateDenyKills += oldStats.GateDenyKills
		newStats.KilledInGate += oldStats.KilledInGate
		newStats.LeftGate += oldStats.LeftGate

		newStats.SpeedDroneUptime += oldStats.SpeedDroneUptime
		newStats.VanillaWarriorUptime += oldStats.VanillaWarriorUptime
		newStats.SpeedWarriorUptime += oldStats.SpeedWarriorUptime

		newStats.GamesPlayed += oldStats.GamesPlayed

		// Merge counters
		for x, row1 := range oldStats.BumpCounter {
			for y, row2 := range row1 {
				newStats.BumpCounter[x][y] += row2
			}
		}
		for x, row1 := range oldStats.KillCounter {
			for y, row2 := range row1 {
				newStats.KillCounter[x][y] += row2
			}
		}
		for x, row1 := range oldStats.DeathCounter {
			for y, row2 := range row1 {
				newStats.DeathCounter[x][y] += row2
			}
		}
	}

	return newStats
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

func (s *PlayerStats) MilKD() float64 {
	return float64(s.MilKills()) / float64(s.MilDeaths())
}

func (s *PlayerStats) QueenKills() int {
	return sumCounterByNestedKey(s.KillCounter, classQueen)
}
