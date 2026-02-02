package state_machine

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
)

func (s *StateMachine) GetOffSnail(event *models.GetOffSnailEvent, hmevent *hivemind.HivemindEvent) {
	rider := s.player(event.Drone)
	rider.OnSnail = false
	s.stats(event.Drone).recordSnailDistance(event.X - rider.LastRecordedSnailX)
	rider.LastRecordedSnailX = event.X
	rider.LastRecordedSnailTime = hmevent.Timestamp
}

func (s *StateMachine) GetOnSnail(event *models.GetOnSnailEvent, hmevent *hivemind.HivemindEvent) {
	player := s.player(event.Drone)
	player.OnSnail = true
	player.LastRecordedSnailX = event.X
	player.LastRecordedSnailTime = hmevent.Timestamp
}

func (s *StateMachine) SnailEat(event *models.SnailEatEvent, hmevent *hivemind.HivemindEvent) {
	s.player(event.Victim).BeingEaten = true

	rider := s.player(event.Rider)
	s.stats(event.Rider).recordSnailDistance(event.X - rider.LastRecordedSnailX)
	rider.LastRecordedSnailX = event.X
	rider.LastRecordedSnailTime = hmevent.Timestamp

	// TODO: Sac / eat stats
}

func (s *StateMachine) SnailEscape(event *models.SnailEscapeEvent) {
	s.player(event.Escapee).BeingEaten = false

	// TODO: Sac stats
}
