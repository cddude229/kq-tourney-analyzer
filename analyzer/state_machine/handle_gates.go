package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) BlessMaiden(event *models.BlessMaidenEvent) {
	// TODO
}

func (s *StateMachine) ReserveMaiden(event *models.ReserveMaidenEvent) {
	// TODO
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
