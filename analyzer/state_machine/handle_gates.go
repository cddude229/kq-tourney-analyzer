package state_machine

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
)

func (s *StateMachine) BlessMaiden(event *models.BlessMaidenEvent, hmevent *hivemind.HivemindEvent) {
	gate := s.Gate(event.X, event.Y)

	timeToAdd := gate.LastTagged.UnixMilli() - hmevent.Timestamp.UnixMilli()
	if gate.Color == nil {
		gate.TimeUntagged += timeToAdd
	} else if *gate.Color == models.GoldTeam1 {
		gate.TimeForGold += timeToAdd
	} else {
		gate.TimeForBlue += timeToAdd
	}

	gate.Color = &event.Team
	gate.LastTagged = hmevent.Timestamp
}

func (s *StateMachine) ReserveMaiden(event *models.ReserveMaidenEvent) {
	// TODO: Derive some interesting stats
}

func (s *StateMachine) UnreserveMaiden(event *models.UnreserveMaidenEvent) {
	if event.Killer != nil {
		s.stats(*event.Killer).GateDenyKills++
		s.stats(event.Drone).KilledInGate++
	} else {
		// TODO: Might have been bumped out?  Compare with glance events
		s.stats(event.Drone).LeftGate++
	}
}

func (s *StateMachine) UseMaiden(event *models.UseMaidenEvent) {
	s.remainingBerries--

	if event.GateType == models.SpeedGate {
		s.player(event.Player).HasSpeed = true
	} else {
		s.player(event.Player).IsWarrior = true
	}
}
