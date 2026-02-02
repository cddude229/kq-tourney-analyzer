package hivemind

func GroupEvents(events []HivemindEvent, matches []TourneyMatch) map[TourneyMatch][]HivemindEvent {
	holder := make(map[TourneyMatch][]HivemindEvent)

	matchesById := make(map[int64]TourneyMatch)
	for _, match := range matches {
		matchesById[match.Id] = match
	}

	for _, event := range events {
		if match, ok := matchesById[event.GameId]; ok {
			holder[match] = append(holder[match], event)
		}
	}

	return holder
}

type MatchesAndStats struct {
	MatchesToEvents map[TourneyMatch][]HivemindEvent
}
