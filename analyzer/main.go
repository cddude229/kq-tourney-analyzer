package main

import (
	"cddude229/kq-tourney-analyzer/aggregation"
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
	"log"
	"sort"
	"time"
)

// TODO: Command line support

func main() {
	// Parsing
	log.Println("Parsing events and matches...")
	parseStart := time.Now()
	events, matches, err := hivemind.OpenAndParseZip("./tourney_data/export_20260127_014624_GDC9.zip")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Parsed %d events and %d matches in %dms",
		len(events), len(matches), time.Now().UnixMilli()-parseStart.UnixMilli())

	// Grouping
	log.Println("Grouping by tourney match...")
	aggStart := time.Now()
	groups := hivemind.GroupEvents(events, matches)
	log.Printf("Grouped in %dms", time.Now().UnixMilli()-aggStart.UnixMilli())

	// Processing
	log.Println("Processing game events...")
	processStart := time.Now()
	stateMachineGroups := make([]aggregation.StateMachineGrouping, 0)

	for match, groupedEvents := range groups {
		sm := models.New()

		for _, event := range groupedEvents {
			smEvent, err := event.ToSMEvent()
			if err != nil {
				log.Fatal(err)
			}

			smEvent.Apply(sm, event.Timestamp)

		}

		group := aggregation.StateMachineGrouping{
			StateMachine: sm,
			TourneyMatch: match,
			GameId:       groupedEvents[0].GameId,
		}

		stateMachineGroups = append(stateMachineGroups, group)
	}

	log.Printf("Processed %d events for %d games in %dms",
		len(events), len(stateMachineGroups), time.Now().UnixMilli()-processStart.UnixMilli())

	// Aggregating
	log.Println("Mapping state machines into players and remapping names...")
	aggStart = time.Now()

	// TODO: Implement remapping functions for tourneys for known cases
	playersAndStats := aggregation.ExtractPlayersForAggregation(stateMachineGroups, []aggregation.PlayerNameGenerator{})

	log.Printf("Mapped state machines into %d users in %dms",
		len(playersAndStats), time.Now().UnixMilli()-aggStart.UnixMilli())

	// And now merging
	log.Println("Merging players for tourney matches...")
	mergeStart := time.Now()

	mergedStats := aggregation.MergeAllPlayersByNameAndTeam(playersAndStats)

	log.Printf("Merged %d players into %d players in %dms",
		len(playersAndStats), len(mergedStats), time.Now().UnixMilli()-mergeStart.UnixMilli())

	// And lastly, stat it
	log.Println("Top 50 mil K/D warriors for this tourney...")
	sort.Slice(mergedStats, func(i, j int) bool {
		return mergedStats[i].PlayerStats.MilKD() > mergedStats[j].PlayerStats.MilKD()
	})

	for _, stats := range mergedStats {
		if !stats.OriginalPlayerId.IsQueen() {
			log.Printf("%s (%s): %f", stats.Name, stats.TeamName, stats.PlayerStats.MilKD())
		}
	}
}
