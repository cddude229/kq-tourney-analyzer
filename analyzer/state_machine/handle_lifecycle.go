package state_machine

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
	"log"
	"math"
)

func (s *StateMachine) GameEnd(event *models.GameEndEvent) {
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

func (s *StateMachine) GameStart(event *models.GameStartEvent) {
	s.mapName = event.MapName
	s.goldOnLeft = event.GoldOnLeft
	s.cabVersion = event.CabVersion
}

func (s *StateMachine) MapStart(event *models.MapStartEvent) {
	s.mapName = event.MapName
	s.goldOnLeft = event.GoldOnLeft
	s.attractMode = event.AttractMode
	s.cabVersion = event.CabVersion
}

func (s *StateMachine) Victory(event *models.VictoryEvent, hmevent *hivemind.HivemindEvent) {
	s.winningTeam = &event.Team
	s.winCondition = &event.WinCondition

	// Add estimated snail distance to the rider(s)
	for playerId, playerState := range s.playerState {
		if playerState.OnSnail && !playerState.IsEating { // LastRecordedSnailX is accurate if we're eating
			pixelsPerSecond := 20.896215463
			if playerState.HasSpeed {
				pixelsPerSecond = 28.209890875
			}

			millisSinceLastUpdate := float64(hmevent.Timestamp.UnixMilli() - playerState.LastRecordedSnailTime.UnixMilli())

			s.stats(playerId).SnailDistance += int(math.Floor(pixelsPerSecond * millisSinceLastUpdate / 1000.0))
		}
	}
}
