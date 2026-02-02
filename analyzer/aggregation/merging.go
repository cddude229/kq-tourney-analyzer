package aggregation

import "cddude229/kq-tourney-analyzer/models"

func MergeAllPlayersByNameAndTeam(players []PlayerAndStats) []PlayerAndStats {
	mapping := map[string][]PlayerAndStats{}
	for _, player := range players {
		mapKey := player.Name + " / " + player.TeamName
		mapping[mapKey] = append(mapping[mapKey], player)
	}

	result := []PlayerAndStats{}

	for _, playersToMerge := range mapping {
		queenFound, droneFound := false, false
		firstPlayer := playersToMerge[0]

		otherStats := []*models.PlayerStats{}
		for _, playerToMerge := range playersToMerge[1:] {
			if playerToMerge.OriginalPlayerId.IsQueen() {
				queenFound = true
			} else {
				droneFound = true
			}
			otherStats = append(otherStats, playerToMerge.PlayerStats)
		}

		playerId := models.PlayerId(0)
		if queenFound && !droneFound {
			playerId = models.BlueQueen
		} else if droneFound && !queenFound {
			playerId = firstPlayer.OriginalPlayerId
		}

		mergedStats := firstPlayer.PlayerStats.Merge(otherStats...)
		result = append(result, PlayerAndStats{
			Name:        firstPlayer.Name,
			TeamName:    firstPlayer.TeamName,
			PlayerStats: mergedStats,

			OriginalPlayerId: playerId,
		})
	}

	return result

}
