package models

import (
	"log"
	"math"
	"time"
)

func (event *GameEndEvent) Apply(s *StateMachine, time time.Time) {
	// WARN: Not always sent in certain builds, including 17.26
	s.mapName = event.MapName
	s.finalGameDuration = event.Duration

	if s.goldOnLeft != event.GoldOnLeft {
		log.Fatalln("goldOnLeft mismatch")
	}

	if s.attractMode != event.AttractMode {
		log.Fatalln("attractMode mismatch")
	}
}

func (event *GameStartEvent) Apply(s *StateMachine, time time.Time) {
	s.mapName = event.MapName
	s.goldOnLeft = event.GoldOnLeft
	s.cabVersion = event.CabVersion
}

func (event *MapStartEvent) Apply(s *StateMachine, time time.Time) {
	s.mapName = event.MapName
	s.goldOnLeft = event.GoldOnLeft
	s.attractMode = event.AttractMode
	s.cabVersion = event.CabVersion
}

func (event *VictoryEvent) Apply(s *StateMachine, time time.Time) {
	s.winningTeam = &event.Team
	s.winCondition = &event.WinCondition

	// Add estimated snail distance to the rider(s)
	// Copy of what I implemented here: https://gitlab.com/kqhivemind/hivemind/-/merge_requests/45/diffs
	for playerId, playerState := range s.playerState {
		if playerState.OnSnail && !playerState.IsEating { // LastRecordedSnailX is accurate if we're eating
			pixelsPerSecond := 20.896215463
			if playerState.HasSpeed {
				pixelsPerSecond = 28.209890875
			}

			millisSinceLastUpdate := float64(time.UnixMilli() - playerState.LastRecordedSnailTime.UnixMilli())

			s.stats(playerId).SnailDistance += int(math.Floor(pixelsPerSecond * millisSinceLastUpdate / 1000.0))
		}
	}
}
