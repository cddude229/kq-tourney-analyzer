package main

import (
	"cddude229/kq-tourney-analyzer/aggregation"
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
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

	gens, err := aggregation.TeamsFromCsvFile("./tourney_data/gdc_players_mapping.csv")
	if err != nil {
		log.Fatal(err)
	}

	playersAndStats := aggregation.ExtractPlayersForAggregation(stateMachineGroups, gens)

	log.Printf("Mapped state machines into %d users in %dms",
		len(playersAndStats), time.Now().UnixMilli()-aggStart.UnixMilli())

	// And now merging
	log.Println("Merging players for tourney matches...")
	mergeStart := time.Now()

	mergedStats := aggregation.MergeAllPlayersByNameAndTeam(playersAndStats)

	log.Printf("Merged %d players into %d players in %dms",
		len(playersAndStats), len(mergedStats), time.Now().UnixMilli()-mergeStart.UnixMilli())

	// And lastly, stat it
	log.Println("Top mil K/D warriors for this tourney...")
	sort.Slice(mergedStats, func(i, j int) bool {
		return mergedStats[i].PlayerStats.MilKD() > mergedStats[j].PlayerStats.MilKD()
	})

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Player", "Team", "G", "K", "D", "K/D", "Mil K", "Mil D", "Mil K/D", "Q Kills", "Q K/Game", "Q K/min", "Mil KPM", "Vanilla KPM", "Speed KPM", "War Time", "%War", "%Speed"})

	for _, stats := range mergedStats {
		if !stats.OriginalPlayerId.IsQueen() {
			t.AppendRow([]interface{}{
				stats.Name,
				stats.TeamName,
				stats.PlayerStats.GamesPlayed,
				stats.PlayerStats.TotalKills(),
				stats.PlayerStats.TotalDeaths(),
				fmt.Sprintf("%.2f", stats.PlayerStats.TotalKD()),
				stats.PlayerStats.MilKills(),
				stats.PlayerStats.MilDeaths(),
				fmt.Sprintf("%.2f", stats.PlayerStats.MilKD()),
				stats.PlayerStats.QueenKills(),
				fmt.Sprintf("%.2f", stats.PlayerStats.QueenKillsPerGame()),
				fmt.Sprintf("%.3f", float64(stats.PlayerStats.QueenKills())*1000.0*60.0/float64(stats.PlayerStats.TotalGameTime)),
				fmt.Sprintf("%.3f", float64(stats.PlayerStats.MilKills())*1000.0*60.0/float64(stats.PlayerStats.VanillaWarriorUptime+stats.PlayerStats.SpeedWarriorUptime)),
				fmt.Sprintf("%.3f", float64(stats.PlayerStats.VanillaMilKills())*1000.0*60.0/float64(stats.PlayerStats.VanillaWarriorUptime)), // TODO: This doesn't match Ryan's yet
				fmt.Sprintf("%.3f", float64(stats.PlayerStats.SpeedMilKills())*1000.0*60.0/float64(stats.PlayerStats.SpeedWarriorUptime)),     // TODO: This doesn't match Ryan's yet
				fmt.Sprintf("%.1f m", stats.PlayerStats.WarriorTimeMinutes()),
				fmt.Sprintf("%.1f%%", float64(stats.PlayerStats.VanillaWarriorUptime+stats.PlayerStats.SpeedWarriorUptime)/float64(stats.PlayerStats.TotalGameTime)*100.0),
				fmt.Sprintf("%.1f%%", float64(stats.PlayerStats.SpeedWarriorUptime)/float64(stats.PlayerStats.VanillaWarriorUptime+stats.PlayerStats.SpeedWarriorUptime)*100.0),
			})
		}
	}

	t.Render()
}
