package models

import (
	"time"
)

func (event *BlessMaidenEvent) Apply(s *StateMachine, time time.Time) {
	gate := s.Gate(event.X, event.Y)

	timeToAdd := gate.LastTagged.UnixMilli() - time.UnixMilli()
	if gate.Color == nil {
		gate.TimeUntagged += timeToAdd
	} else if *gate.Color == GoldTeam1 {
		gate.TimeForGold += timeToAdd
	} else {
		gate.TimeForBlue += timeToAdd
	}

	gate.Color = &event.Team
	gate.LastTagged = time
}

func (event *ReserveMaidenEvent) Apply(s *StateMachine, time time.Time) {}

func (event *UnreserveMaidenEvent) Apply(s *StateMachine, time time.Time) {
	if event.Killer != nil {
		s.stats(*event.Killer).GateDenyKills++
		s.stats(event.Drone).KilledInGate++
	} else {
		// TODO: Might have been bumped out?  Compare with glance events
		s.stats(event.Drone).LeftGate++
	}
}

func (event *UseMaidenEvent) Apply(s *StateMachine, time time.Time) {
	s.remainingBerries--

	if event.GateType == SpeedGate {
		s.player(event.Player).HasSpeed = true
	} else {
		s.player(event.Player).IsWarrior = true
	}
}
