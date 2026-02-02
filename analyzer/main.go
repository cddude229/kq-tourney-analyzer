package main

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/state_machine"
	"log"
	"time"
)

// TODO: Command line support

func main() {

	log.Println("Parsing events...")

	parseStart := time.Now()
	events, err := hivemind.OpenAndParseZip("./tourney_data/export_20260127_014624_GDC9.zip")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Parsed %d events in %dms",
		len(events),
		time.Now().UnixMilli()-parseStart.UnixMilli())

	// Sanity check the events are sorted in order
	//for _, event := range events[0:10] {
	//	log.Printf("timestamp: %s", event.Timestamp)
	//}

	processStart := time.Now()
	stateMachineMap := make(map[int64]*state_machine.StateMachine)
	skippedEvents := 0

	log.Println("Processing events...")
	for _, event := range events {
		sm, exists := stateMachineMap[event.GameId]
		if !exists {
			sm = state_machine.New()
			stateMachineMap[event.GameId] = sm
		}

		processed, err := sm.HandleHivemindEvent(event)
		if err != nil {
			log.Fatal(err)
		}

		if !processed {
			skippedEvents++
		}
	}

	log.Printf("Processed %d events for %d games in %dms",
		len(events)-skippedEvents,
		len(stateMachineMap),
		time.Now().UnixMilli()-processStart.UnixMilli())
}
