package models

import (
	"time"
)

func (event *BlessMaidenEvent) Apply(s *StateMachine, time time.Time) {
	gate := s.Gate(event.X, event.Y)

	// Don't count time before the map started
	lastTagged := gate.LastTagged
	if s.startTime != nil && s.startTime.After(gate.LastTagged) {
		lastTagged = *s.startTime
	}

	if gate.Color == nil {
		// Do nothing for now
	} else if *gate.Color == GoldTeam1 {
		gate.TimeForGold += time.UnixMilli() - lastTagged.UnixMilli()
	} else if *gate.Color == BlueTeam1 {
		gate.TimeForBlue += time.UnixMilli() - lastTagged.UnixMilli()
	}

	gate.LastTagged = time
	gate.Color = &event.Team
}

func (event *ReserveMaidenEvent) Apply(s *StateMachine, time time.Time) {}

func (event *UnreserveMaidenEvent) Apply(s *StateMachine, time time.Time) {
	if event.Killer != nil {
		s.Stats(*event.Killer).GateDenyKills++
		s.Stats(event.Drone).KilledInGate++
	} else {
		// TODO: Might have been bumped out?  Compare with glance events
		s.Stats(event.Drone).LeftGate++
	}
}

func (event *UseMaidenEvent) Apply(s *StateMachine, time time.Time) {
	s.remainingBerries--

	player := s.player(event.Player)
	gate := s.Gate(event.X, event.Y)

	gate.gateType = &event.GateType
	gate.TimesUsed++

	if event.GateType == SpeedGate {
		player.IsSpeed = true
		player.GotSpeedAt = time
	} else {
		player.IsWarrior = true
		player.BecameWarriorAt = time

		if player.IsSpeed {
			s.Stats(event.Player).SpeedDroneUptime += time.UnixMilli() - player.GotSpeedAt.UnixMilli()
		}
	}
}
