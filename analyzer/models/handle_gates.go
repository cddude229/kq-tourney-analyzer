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

	player := s.player(event.Player)
	if event.GateType == SpeedGate {
		player.IsSpeed = true
		player.GotSpeedAt = time
	} else {
		player.IsWarrior = true
		player.BecameWarriorAt = time

		if player.IsSpeed {
			s.stats(event.Player).SpeedDroneUptime += time.UnixMilli() - player.GotSpeedAt.UnixMilli()
		}
	}
}
