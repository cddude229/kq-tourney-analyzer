package aggregation

import (
	"cddude229/kq-tourney-analyzer/models"
	"fmt"
	"log"
)

// A PlayerNameGenerator converts a cab spot and other metadata into some groupable name to represent that person
type PlayerNameGenerator interface {
	// NamePlayer either returns the name to use, or Nil to mean "does not apply here"
	NamePlayer(id models.PlayerId, gameId int64, matchId int64, teamName string) *string
}

func ExtractPlayersForAggregation(groups []StateMachineGrouping, playerNameGenerators []PlayerNameGenerator) []PlayerAndStats {
	playerAndStats := make([]PlayerAndStats, 0)

	for _, group := range groups {
		for id := 1; id <= 10; id++ {
			playerId := models.PlayerId(id)
			gameId := group.GameId
			matchId := group.TourneyMatch.TournamentMatchId

			teamName := group.TourneyMatch.GoldTeamName
			if playerId.IsBlue() {
				teamName = group.TourneyMatch.BlueTeamName
			}

			playerName := defaultPlayerName(playerId)
			for _, playerNameGenerator := range playerNameGenerators {
				potentialName := playerNameGenerator.NamePlayer(playerId, gameId, matchId, teamName)
				if potentialName != nil {
					playerName = *potentialName
					break
				}
			}

			playerAndStats = append(playerAndStats, PlayerAndStats{
				Name:             playerName,
				OriginalPlayerId: playerId,

				GameId:   gameId,
				MatchId:  matchId,
				TeamName: teamName,

				PlayerStats: group.StateMachine.Stats(playerId),
			})
		}
	}

	return playerAndStats
}

type PlayerAndStats struct {
	Name             string
	OriginalPlayerId models.PlayerId

	GameId   int64
	MatchId  int64
	TeamName string

	PlayerStats *models.PlayerStats
}

func defaultPlayerName(id models.PlayerId) string {
	switch id {
	case models.GoldStripes, models.BlueChex:
		return "Gold Stripes / Blue Chex"
	case models.GoldAbs, models.BlueSkulls:
		return "Gold Abs / Blue Skulls"
	case models.GoldQueen, models.BlueQueen:
		return "Gold Queen / Blue Queen"
	case models.GoldSkulls, models.BlueAbs:
		return "Gold Skulls / Blue Abs"
	case models.GoldChex, models.BlueStripes:
		return "Gold Chex / Blue Strioes"
	}

	log.Fatalln(fmt.Sprintf("unknown player ID: %d", id))
	return ""
}
