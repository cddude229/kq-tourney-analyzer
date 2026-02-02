package models

import (
	"time"
)

func (event *GetOffSnailEvent) Apply(s *StateMachine, time time.Time) {
	rider := s.player(event.Drone)
	rider.OnSnail = false
	s.stats(event.Drone).recordSnailDistance(event.X - rider.LastRecordedSnailX)
	rider.LastRecordedSnailX = event.X
	rider.LastRecordedSnailTime = time
}

func (event *GetOnSnailEvent) Apply(s *StateMachine, time time.Time) {
	player := s.player(event.Drone)
	player.OnSnail = true
	player.LastRecordedSnailX = event.X
	player.LastRecordedSnailTime = time
}

func (event *SnailEatEvent) Apply(s *StateMachine, time time.Time) {
	s.player(event.Victim).BeingEaten = true

	rider := s.player(event.Rider)
	s.stats(event.Rider).recordSnailDistance(event.X - rider.LastRecordedSnailX)
	rider.LastRecordedSnailX = event.X
	rider.LastRecordedSnailTime = time

	// TODO: Sac / eat stats
}

func (event *SnailEscapeEvent) Apply(s *StateMachine, time time.Time) {
	s.player(event.Escapee).BeingEaten = false

	// TODO: Sac stats
}
