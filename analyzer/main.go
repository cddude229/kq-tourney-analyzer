package main

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"log"
)

// TODO: Command line support

func main() {
	events, err := hivemind.OpenAndParseZip("./tourney_data/export_20260127_014624.zip")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d parsed events", len(events))
}
