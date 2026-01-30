package main

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/state_machine"
	"log"
)

// TODO: Command line support

func main() {

	log.Println("Parsing events...")

	events, err := hivemind.OpenAndParseZip("./tourney_data/export_20260127_014624.zip")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Parsed %d events", len(events))

	// Sanity check the events are sorted in order
	//for _, event := range events[0:10] {
	//	log.Printf("timestamp: %s", event.Timestamp)
	//}

	stateMachineMap := make(map[int64]state_machine.StateMachine)
	skippedEvents := 0

	log.Println("Processing events...")
	for _, event := range events {
		sm, exists := stateMachineMap[event.GameId]
		if !exists {
			sm = state_machine.StateMachine{}
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

	log.Printf("Processed %d events for %d games", len(events)-skippedEvents, len(stateMachineMap))
}
