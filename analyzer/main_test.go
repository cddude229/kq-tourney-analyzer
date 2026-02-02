package main

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame1652756Accuracy(t *testing.T) {
	// One game from GDC - https://kqhivemind.com/game/1652756 and https://kqhivemind.com/api/game/game/1652756/stats/
	events, _, err := hivemind.OpenAndParseZip("./test_data/game_1652756.zip")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, events)

	// Process!
	sm := models.New()
	for _, event := range events {
		smevent, err := event.ToSMEvent()
		assert.Nil(t, err)
		smevent.Apply(sm, event.Timestamp)
	}

	t.Run("Total kills", func(t *testing.T) {
		assert.Equal(t, 10, sm.Stats(models.BlueStripes).TotalKills())
		assert.Equal(t, 6, sm.Stats(models.BlueAbs).TotalKills())
		assert.Equal(t, 7, sm.Stats(models.BlueQueen).TotalKills())
		assert.Equal(t, 2, sm.Stats(models.BlueSkulls).TotalKills())
		assert.Equal(t, 4, sm.Stats(models.BlueChex).TotalKills())

		assert.Equal(t, 5, sm.Stats(models.GoldStripes).TotalKills())
		assert.Equal(t, 3, sm.Stats(models.GoldAbs).TotalKills())
		assert.Equal(t, 8, sm.Stats(models.GoldQueen).TotalKills())
		assert.Equal(t, 6, sm.Stats(models.GoldSkulls).TotalKills())
		assert.Equal(t, 0, sm.Stats(models.GoldChex).TotalKills())
	})

	t.Run("Total deaths", func(t *testing.T) {
		assert.Equal(t, 6, sm.Stats(models.BlueStripes).TotalDeaths())
		assert.Equal(t, 8, sm.Stats(models.BlueAbs).TotalDeaths())
		assert.Equal(t, 0, sm.Stats(models.BlueQueen).TotalDeaths())
		assert.Equal(t, 3, sm.Stats(models.BlueSkulls).TotalDeaths())
		assert.Equal(t, 5, sm.Stats(models.BlueChex).TotalDeaths())

		assert.Equal(t, 8, sm.Stats(models.GoldStripes).TotalDeaths())
		assert.Equal(t, 5, sm.Stats(models.GoldAbs).TotalDeaths())
		assert.Equal(t, 3, sm.Stats(models.GoldQueen).TotalDeaths())
		assert.Equal(t, 3, sm.Stats(models.GoldSkulls).TotalDeaths())
		assert.Equal(t, 10, sm.Stats(models.GoldChex).TotalDeaths())
	})

	t.Run("Mil kills", func(t *testing.T) {
		assert.Equal(t, 5, sm.Stats(models.BlueStripes).MilKills())
		assert.Equal(t, 2, sm.Stats(models.BlueAbs).MilKills())
		assert.Equal(t, 5, sm.Stats(models.BlueQueen).MilKills())
		assert.Equal(t, 0, sm.Stats(models.BlueSkulls).MilKills())
		assert.Equal(t, 1, sm.Stats(models.BlueChex).MilKills())

		assert.Equal(t, 1, sm.Stats(models.GoldStripes).MilKills())
		assert.Equal(t, 3, sm.Stats(models.GoldAbs).MilKills())
		assert.Equal(t, 7, sm.Stats(models.GoldQueen).MilKills())
		assert.Equal(t, 2, sm.Stats(models.GoldSkulls).MilKills())
		assert.Equal(t, 0, sm.Stats(models.GoldChex).MilKills())
	})

	t.Run("Mil deaths", func(t *testing.T) {
		assert.Equal(t, 4, sm.Stats(models.BlueStripes).MilDeaths())
		assert.Equal(t, 5, sm.Stats(models.BlueAbs).MilDeaths())
		assert.Equal(t, 0, sm.Stats(models.BlueQueen).MilDeaths())
		assert.Equal(t, 2, sm.Stats(models.BlueSkulls).MilDeaths())
		assert.Equal(t, 2, sm.Stats(models.BlueChex).MilDeaths())

		assert.Equal(t, 5, sm.Stats(models.GoldStripes).MilDeaths())
		assert.Equal(t, 3, sm.Stats(models.GoldAbs).MilDeaths())
		assert.Equal(t, 3, sm.Stats(models.GoldQueen).MilDeaths())
		assert.Equal(t, 2, sm.Stats(models.GoldSkulls).MilDeaths())
		assert.Equal(t, 0, sm.Stats(models.GoldChex).MilDeaths())
	})

	t.Run("Queen kills", func(t *testing.T) {
		assert.Equal(t, 1, sm.Stats(models.BlueStripes).QueenKills())
		assert.Equal(t, 0, sm.Stats(models.BlueAbs).QueenKills())
		assert.Equal(t, 2, sm.Stats(models.BlueQueen).QueenKills())
		assert.Equal(t, 0, sm.Stats(models.BlueSkulls).QueenKills())
		assert.Equal(t, 0, sm.Stats(models.BlueChex).QueenKills())

		assert.Equal(t, 0, sm.Stats(models.GoldStripes).QueenKills())
		assert.Equal(t, 0, sm.Stats(models.GoldAbs).QueenKills())
		assert.Equal(t, 0, sm.Stats(models.GoldQueen).QueenKills())
		assert.Equal(t, 0, sm.Stats(models.GoldSkulls).QueenKills())
		assert.Equal(t, 0, sm.Stats(models.GoldChex).QueenKills())
	})

	t.Run("Berries", func(t *testing.T) {
		assert.Equal(t, 0, sm.Stats(models.BlueStripes).TotalBerries())
		assert.Equal(t, 0, sm.Stats(models.BlueAbs).TotalBerries())
		assert.Equal(t, 0, sm.Stats(models.BlueQueen).TotalBerries())
		assert.Equal(t, 0, sm.Stats(models.BlueSkulls).TotalBerries())
		assert.Equal(t, 0, sm.Stats(models.BlueChex).TotalBerries())

		assert.Equal(t, 0, sm.Stats(models.GoldStripes).TotalBerries())
		assert.Equal(t, 0, sm.Stats(models.GoldAbs).TotalBerries())
		assert.Equal(t, 0, sm.Stats(models.GoldQueen).TotalBerries())
		assert.Equal(t, 0, sm.Stats(models.GoldSkulls).TotalBerries())
		assert.Equal(t, 0, sm.Stats(models.GoldChex).TotalBerries())
	})

	t.Run("Snail distance", func(t *testing.T) {
		assert.Equal(t, 346, sm.Stats(models.BlueStripes).SnailDistance)
		assert.Equal(t, 504, sm.Stats(models.BlueAbs).SnailDistance)
		assert.Equal(t, 0, sm.Stats(models.BlueQueen).SnailDistance)
		assert.Equal(t, 0, sm.Stats(models.BlueSkulls).SnailDistance)
		assert.Equal(t, 114, sm.Stats(models.BlueChex).SnailDistance)

		assert.Equal(t, 0, sm.Stats(models.GoldStripes).SnailDistance)
		assert.Equal(t, 0, sm.Stats(models.GoldAbs).SnailDistance)
		assert.Equal(t, 0, sm.Stats(models.GoldQueen).SnailDistance)
		assert.Equal(t, 0, sm.Stats(models.GoldSkulls).SnailDistance)
		assert.Equal(t, 1143, sm.Stats(models.GoldChex).SnailDistance)
	})

	t.Run("Vanilla Warrior Uptime", func(t *testing.T) {
		assert.Equal(t, int64(53669), sm.Stats(models.BlueStripes).VanillaWarriorUptime)
		assert.Equal(t, int64(39484), sm.Stats(models.BlueAbs).VanillaWarriorUptime)
		assert.Equal(t, int64(47344), sm.Stats(models.BlueSkulls).VanillaWarriorUptime)
		assert.Equal(t, int64(3533), sm.Stats(models.BlueChex).VanillaWarriorUptime)

		assert.Equal(t, int64(50237), sm.Stats(models.GoldStripes).VanillaWarriorUptime)
		assert.Equal(t, int64(0), sm.Stats(models.GoldAbs).VanillaWarriorUptime)
		assert.Equal(t, int64(0), sm.Stats(models.GoldSkulls).VanillaWarriorUptime)
		assert.Equal(t, int64(0), sm.Stats(models.GoldChex).VanillaWarriorUptime)
	})

	t.Run("Speed Warrior Uptime", func(t *testing.T) {
		assert.Equal(t, int64(26246), sm.Stats(models.BlueStripes).SpeedWarriorUptime)
		assert.Equal(t, int64(47501), sm.Stats(models.BlueAbs).SpeedWarriorUptime)
		assert.Equal(t, int64(103524), sm.Stats(models.BlueSkulls).SpeedWarriorUptime)
		assert.Equal(t, int64(93689), sm.Stats(models.BlueChex).SpeedWarriorUptime)

		assert.Equal(t, int64(23627), sm.Stats(models.GoldStripes).SpeedWarriorUptime)
		assert.Equal(t, int64(90827), sm.Stats(models.GoldAbs).SpeedWarriorUptime)
		assert.Equal(t, int64(156920), sm.Stats(models.GoldSkulls).SpeedWarriorUptime)
		assert.Equal(t, int64(0), sm.Stats(models.GoldChex).SpeedWarriorUptime)
	})

	t.Run("Total gate control", func(t *testing.T) {
		// TODO: Waiting on Abby to get back to me about a bug in hivemind API.  Her numbers don't include gates that
		// just were never used
		// _, _, _ := sm.CalculateGateControlTimeInSeconds()

		// assert.Equal(t, 213.284+187.367, blue)
		// assert.Equal(t, 184.191+224.913, gold)
	})

	t.Run("Warrior gate control", func(t *testing.T) {
		// TODO: Waiting on Abby to get back to me about a bug in hivemind API.  Her numbers don't include gates that
		// just were never used
		// _, _, _ := sm.CalculateWarriorGateControlTimeInSeconds()

		// assert.Equal(t, 213.284+187.367, blue)
		// assert.Equal(t, 184.191+224.913, gold)
	})

	t.Run("Speed gate control", func(t *testing.T) {
		blue, gold := sm.CalculateSpeedGateControlTimeInSeconds()

		assert.Equal(t, 187.367, blue)
		assert.Equal(t, 224.913, gold)
	})

}
