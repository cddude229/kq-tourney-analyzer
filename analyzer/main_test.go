package main

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
	"cddude229/kq-tourney-analyzer/state_machine"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame1652756Accuracy(t *testing.T) {
	// One game from GDC - https://kqhivemind.com/game/1652756 and https://kqhivemind.com/api/game/game/1652756/stats/
	events, err := hivemind.OpenAndParseZip("./test_data/game_1652756.zip")
	assert.Nil(t, err)
	assert.NotEmpty(t, events)

	// Process!
	sm := state_machine.New()
	for _, event := range events {
		_, err = sm.HandleHivemindEvent(event)
		assert.Nil(t, err)
	}

	// Now validate all the stats...
	// Total kills
	assert.Equal(t, sm.PlayerStats[models.BlueStripes].TotalKills(), 10)
	assert.Equal(t, sm.PlayerStats[models.BlueAbs].TotalKills(), 6)
	assert.Equal(t, sm.PlayerStats[models.BlueQueen].TotalKills(), 7)
	assert.Equal(t, sm.PlayerStats[models.BlueSkulls].TotalKills(), 2)
	assert.Equal(t, sm.PlayerStats[models.BlueChex].TotalKills(), 4)

	assert.Equal(t, sm.PlayerStats[models.GoldStripes].TotalKills(), 5)
	assert.Equal(t, sm.PlayerStats[models.GoldAbs].TotalKills(), 3)
	assert.Equal(t, sm.PlayerStats[models.GoldQueen].TotalKills(), 8)
	assert.Equal(t, sm.PlayerStats[models.GoldSkulls].TotalKills(), 6)
	assert.Equal(t, sm.PlayerStats[models.GoldChex].TotalKills(), 0)

	// Total deaths
	assert.Equal(t, sm.PlayerStats[models.BlueStripes].TotalDeaths(), 6)
	assert.Equal(t, sm.PlayerStats[models.BlueAbs].TotalDeaths(), 8)
	assert.Equal(t, sm.PlayerStats[models.BlueQueen].TotalDeaths(), 0)
	assert.Equal(t, sm.PlayerStats[models.BlueSkulls].TotalDeaths(), 3)
	assert.Equal(t, sm.PlayerStats[models.BlueChex].TotalDeaths(), 5)

	assert.Equal(t, sm.PlayerStats[models.GoldStripes].TotalDeaths(), 8)
	assert.Equal(t, sm.PlayerStats[models.GoldAbs].TotalDeaths(), 5)
	assert.Equal(t, sm.PlayerStats[models.GoldQueen].TotalDeaths(), 3)
	assert.Equal(t, sm.PlayerStats[models.GoldSkulls].TotalDeaths(), 3)
	assert.Equal(t, sm.PlayerStats[models.GoldChex].TotalDeaths(), 10)

	// Mil kills
	assert.Equal(t, sm.PlayerStats[models.BlueStripes].MilKills(), 5)
	assert.Equal(t, sm.PlayerStats[models.BlueAbs].MilKills(), 2)
	assert.Equal(t, sm.PlayerStats[models.BlueQueen].MilKills(), 5)
	assert.Equal(t, sm.PlayerStats[models.BlueSkulls].MilKills(), 0)
	assert.Equal(t, sm.PlayerStats[models.BlueChex].MilKills(), 1)

	assert.Equal(t, sm.PlayerStats[models.GoldStripes].MilKills(), 1)
	assert.Equal(t, sm.PlayerStats[models.GoldAbs].MilKills(), 3)
	assert.Equal(t, sm.PlayerStats[models.GoldQueen].MilKills(), 7)
	assert.Equal(t, sm.PlayerStats[models.GoldSkulls].MilKills(), 2)
	assert.Equal(t, sm.PlayerStats[models.GoldChex].MilKills(), 0)

	// Mil deaths
	assert.Equal(t, sm.PlayerStats[models.BlueStripes].MilDeaths(), 4)
	assert.Equal(t, sm.PlayerStats[models.BlueAbs].MilDeaths(), 5)
	assert.Equal(t, sm.PlayerStats[models.BlueQueen].MilDeaths(), 0)
	assert.Equal(t, sm.PlayerStats[models.BlueSkulls].MilDeaths(), 2)
	assert.Equal(t, sm.PlayerStats[models.BlueChex].MilDeaths(), 2)

	assert.Equal(t, sm.PlayerStats[models.GoldStripes].MilDeaths(), 5)
	assert.Equal(t, sm.PlayerStats[models.GoldAbs].MilDeaths(), 3)
	assert.Equal(t, sm.PlayerStats[models.GoldQueen].MilDeaths(), 3)
	assert.Equal(t, sm.PlayerStats[models.GoldSkulls].MilDeaths(), 2)
	assert.Equal(t, sm.PlayerStats[models.GoldChex].MilDeaths(), 0)

	// Queen kills
	assert.Equal(t, sm.PlayerStats[models.BlueStripes].QueenKills(), 1)
	assert.Equal(t, sm.PlayerStats[models.BlueAbs].QueenKills(), 0)
	assert.Equal(t, sm.PlayerStats[models.BlueQueen].QueenKills(), 2)
	assert.Equal(t, sm.PlayerStats[models.BlueSkulls].QueenKills(), 0)
	assert.Equal(t, sm.PlayerStats[models.BlueChex].QueenKills(), 0)

	assert.Equal(t, sm.PlayerStats[models.GoldStripes].QueenKills(), 0)
	assert.Equal(t, sm.PlayerStats[models.GoldAbs].QueenKills(), 0)
	assert.Equal(t, sm.PlayerStats[models.GoldQueen].QueenKills(), 0)
	assert.Equal(t, sm.PlayerStats[models.GoldSkulls].QueenKills(), 0)
	assert.Equal(t, sm.PlayerStats[models.GoldChex].QueenKills(), 0)

	// Berries.  All 0s though
	assert.Equal(t, sm.PlayerStats[models.BlueStripes].TotalBerries(), 0)
	assert.Equal(t, sm.PlayerStats[models.BlueAbs].TotalBerries(), 0)
	assert.Equal(t, sm.PlayerStats[models.BlueQueen].TotalBerries(), 0)
	assert.Equal(t, sm.PlayerStats[models.BlueSkulls].TotalBerries(), 0)
	assert.Equal(t, sm.PlayerStats[models.BlueChex].TotalBerries(), 0)

	assert.Equal(t, sm.PlayerStats[models.GoldStripes].TotalBerries(), 0)
	assert.Equal(t, sm.PlayerStats[models.GoldAbs].TotalBerries(), 0)
	assert.Equal(t, sm.PlayerStats[models.GoldQueen].TotalBerries(), 0)
	assert.Equal(t, sm.PlayerStats[models.GoldSkulls].TotalBerries(), 0)
	assert.Equal(t, sm.PlayerStats[models.GoldChex].TotalBerries(), 0)

	// Snail distance
	assert.Equal(t, sm.PlayerStats[models.BlueStripes].SnailDistance, 346)
	assert.Equal(t, sm.PlayerStats[models.BlueAbs].SnailDistance, 504)
	assert.Equal(t, sm.PlayerStats[models.BlueQueen].SnailDistance, 0)
	assert.Equal(t, sm.PlayerStats[models.BlueSkulls].SnailDistance, 0)
	assert.Equal(t, sm.PlayerStats[models.BlueChex].SnailDistance, 114)

	assert.Equal(t, sm.PlayerStats[models.GoldStripes].SnailDistance, 0)
	assert.Equal(t, sm.PlayerStats[models.GoldAbs].SnailDistance, 0)
	assert.Equal(t, sm.PlayerStats[models.GoldQueen].SnailDistance, 0)
	assert.Equal(t, sm.PlayerStats[models.GoldSkulls].SnailDistance, 0)
	assert.Equal(t, sm.PlayerStats[models.GoldChex].SnailDistance, 1143)

}
