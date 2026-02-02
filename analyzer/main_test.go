package main

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame1652756Accuracy(t *testing.T) {
	// One game from GDC - https://kqhivemind.com/game/1652756 and https://kqhivemind.com/api/game/game/1652756/stats/
	events, err := hivemind.OpenAndParseZip("./test_data/game_1652756.zip")
	assert.Nil(t, err)
	assert.NotEmpty(t, events)

	// Process!
	sm := models.New()
	for _, event := range events {
		smevent, err := event.ToSMEvent()
		assert.Nil(t, err)
		smevent.Apply(sm, event.Timestamp)
	}

	t.Run("Total kills", func(t *testing.T) {
		assert.Equal(t, 10, sm.PlayerStats[models.BlueStripes].TotalKills())
		assert.Equal(t, 6, sm.PlayerStats[models.BlueAbs].TotalKills())
		assert.Equal(t, 7, sm.PlayerStats[models.BlueQueen].TotalKills())
		assert.Equal(t, 2, sm.PlayerStats[models.BlueSkulls].TotalKills())
		assert.Equal(t, 4, sm.PlayerStats[models.BlueChex].TotalKills())

		assert.Equal(t, 5, sm.PlayerStats[models.GoldStripes].TotalKills())
		assert.Equal(t, 3, sm.PlayerStats[models.GoldAbs].TotalKills())
		assert.Equal(t, 8, sm.PlayerStats[models.GoldQueen].TotalKills())
		assert.Equal(t, 6, sm.PlayerStats[models.GoldSkulls].TotalKills())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldChex].TotalKills())
	})

	t.Run("Total deaths", func(t *testing.T) {
		assert.Equal(t, 6, sm.PlayerStats[models.BlueStripes].TotalDeaths())
		assert.Equal(t, 8, sm.PlayerStats[models.BlueAbs].TotalDeaths())
		assert.Equal(t, 0, sm.PlayerStats[models.BlueQueen].TotalDeaths())
		assert.Equal(t, 3, sm.PlayerStats[models.BlueSkulls].TotalDeaths())
		assert.Equal(t, 5, sm.PlayerStats[models.BlueChex].TotalDeaths())

		assert.Equal(t, 8, sm.PlayerStats[models.GoldStripes].TotalDeaths())
		assert.Equal(t, 5, sm.PlayerStats[models.GoldAbs].TotalDeaths())
		assert.Equal(t, 3, sm.PlayerStats[models.GoldQueen].TotalDeaths())
		assert.Equal(t, 3, sm.PlayerStats[models.GoldSkulls].TotalDeaths())
		assert.Equal(t, 10, sm.PlayerStats[models.GoldChex].TotalDeaths())
	})

	t.Run("Mil kills", func(t *testing.T) {
		assert.Equal(t, 5, sm.PlayerStats[models.BlueStripes].MilKills())
		assert.Equal(t, 2, sm.PlayerStats[models.BlueAbs].MilKills())
		assert.Equal(t, 5, sm.PlayerStats[models.BlueQueen].MilKills())
		assert.Equal(t, 0, sm.PlayerStats[models.BlueSkulls].MilKills())
		assert.Equal(t, 1, sm.PlayerStats[models.BlueChex].MilKills())

		assert.Equal(t, 1, sm.PlayerStats[models.GoldStripes].MilKills())
		assert.Equal(t, 3, sm.PlayerStats[models.GoldAbs].MilKills())
		assert.Equal(t, 7, sm.PlayerStats[models.GoldQueen].MilKills())
		assert.Equal(t, 2, sm.PlayerStats[models.GoldSkulls].MilKills())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldChex].MilKills())
	})

	t.Run("Mil deaths", func(t *testing.T) {
		assert.Equal(t, 4, sm.PlayerStats[models.BlueStripes].MilDeaths())
		assert.Equal(t, 5, sm.PlayerStats[models.BlueAbs].MilDeaths())
		assert.Equal(t, 0, sm.PlayerStats[models.BlueQueen].MilDeaths())
		assert.Equal(t, 2, sm.PlayerStats[models.BlueSkulls].MilDeaths())
		assert.Equal(t, 2, sm.PlayerStats[models.BlueChex].MilDeaths())

		assert.Equal(t, 5, sm.PlayerStats[models.GoldStripes].MilDeaths())
		assert.Equal(t, 3, sm.PlayerStats[models.GoldAbs].MilDeaths())
		assert.Equal(t, 3, sm.PlayerStats[models.GoldQueen].MilDeaths())
		assert.Equal(t, 2, sm.PlayerStats[models.GoldSkulls].MilDeaths())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldChex].MilDeaths())
	})

	t.Run("Queen kills", func(t *testing.T) {
		assert.Equal(t, 1, sm.PlayerStats[models.BlueStripes].QueenKills())
		assert.Equal(t, 0, sm.PlayerStats[models.BlueAbs].QueenKills())
		assert.Equal(t, 2, sm.PlayerStats[models.BlueQueen].QueenKills())
		assert.Equal(t, 0, sm.PlayerStats[models.BlueSkulls].QueenKills())
		assert.Equal(t, 0, sm.PlayerStats[models.BlueChex].QueenKills())

		assert.Equal(t, 0, sm.PlayerStats[models.GoldStripes].QueenKills())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldAbs].QueenKills())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldQueen].QueenKills())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldSkulls].QueenKills())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldChex].QueenKills())
	})

	t.Run("Berries", func(t *testing.T) {
		assert.Equal(t, 0, sm.PlayerStats[models.BlueStripes].TotalBerries())
		assert.Equal(t, 0, sm.PlayerStats[models.BlueAbs].TotalBerries())
		assert.Equal(t, 0, sm.PlayerStats[models.BlueQueen].TotalBerries())
		assert.Equal(t, 0, sm.PlayerStats[models.BlueSkulls].TotalBerries())
		assert.Equal(t, 0, sm.PlayerStats[models.BlueChex].TotalBerries())

		assert.Equal(t, 0, sm.PlayerStats[models.GoldStripes].TotalBerries())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldAbs].TotalBerries())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldQueen].TotalBerries())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldSkulls].TotalBerries())
		assert.Equal(t, 0, sm.PlayerStats[models.GoldChex].TotalBerries())
	})

	t.Run("Snail distance", func(t *testing.T) {
		assert.Equal(t, 346, sm.PlayerStats[models.BlueStripes].SnailDistance)
		assert.Equal(t, 504, sm.PlayerStats[models.BlueAbs].SnailDistance)
		assert.Equal(t, 0, sm.PlayerStats[models.BlueQueen].SnailDistance)
		assert.Equal(t, 0, sm.PlayerStats[models.BlueSkulls].SnailDistance)
		assert.Equal(t, 114, sm.PlayerStats[models.BlueChex].SnailDistance)

		assert.Equal(t, 0, sm.PlayerStats[models.GoldStripes].SnailDistance)
		assert.Equal(t, 0, sm.PlayerStats[models.GoldAbs].SnailDistance)
		assert.Equal(t, 0, sm.PlayerStats[models.GoldQueen].SnailDistance)
		assert.Equal(t, 0, sm.PlayerStats[models.GoldSkulls].SnailDistance)
		assert.Equal(t, 1143, sm.PlayerStats[models.GoldChex].SnailDistance)
	})

	t.Run("Vanilla Warrior Uptime", func(t *testing.T) {
		assert.Equal(t, int64(53669), sm.PlayerStats[models.BlueStripes].VanillaWarriorUptime)
		assert.Equal(t, int64(39484), sm.PlayerStats[models.BlueAbs].VanillaWarriorUptime)
		assert.Equal(t, int64(47344), sm.PlayerStats[models.BlueSkulls].VanillaWarriorUptime)
		assert.Equal(t, int64(3533), sm.PlayerStats[models.BlueChex].VanillaWarriorUptime)

		assert.Equal(t, int64(50237), sm.PlayerStats[models.GoldStripes].VanillaWarriorUptime)
		assert.Equal(t, int64(0), sm.PlayerStats[models.GoldAbs].VanillaWarriorUptime)
		assert.Equal(t, int64(0), sm.PlayerStats[models.GoldSkulls].VanillaWarriorUptime)
		assert.Equal(t, int64(0), sm.PlayerStats[models.GoldChex].VanillaWarriorUptime)
	})

	t.Run("Speed Warrior Uptime", func(t *testing.T) {
		assert.Equal(t, int64(26246), sm.PlayerStats[models.BlueStripes].SpeedWarriorUptime)
		assert.Equal(t, int64(47501), sm.PlayerStats[models.BlueAbs].SpeedWarriorUptime)
		assert.Equal(t, int64(103524), sm.PlayerStats[models.BlueSkulls].SpeedWarriorUptime)
		assert.Equal(t, int64(93689), sm.PlayerStats[models.BlueChex].SpeedWarriorUptime)

		assert.Equal(t, int64(23627), sm.PlayerStats[models.GoldStripes].SpeedWarriorUptime)
		assert.Equal(t, int64(90827), sm.PlayerStats[models.GoldAbs].SpeedWarriorUptime)
		assert.Equal(t, int64(156920), sm.PlayerStats[models.GoldSkulls].SpeedWarriorUptime)
		assert.Equal(t, int64(0), sm.PlayerStats[models.GoldChex].SpeedWarriorUptime)
	})
}
