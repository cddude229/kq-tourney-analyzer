package models

import "time"

func (s *StateMachine) countBerryForTeam(color TeamColor2) {
	if color == BlueTeam2 {
		s.blueBerries++
	} else {
		s.goldBerries++
	}
}

func (event *BerryDepositEvent) Apply(s *StateMachine, time time.Time) {
	s.remainingBerries--

	s.player(event.Player).HasBerry = false

	s.countBerryForTeam(event.Player.Team())
	s.Stats(event.Player).BerriesDunked++
}

func (event *BerryKickInEvent) Apply(s *StateMachine, time time.Time) {
	s.remainingBerries--

	if event.PlayersHive {
		s.countBerryForTeam(event.Player.Team())
		s.Stats(event.Player).BerriesKickedOurTeam++
	} else {
		s.countBerryForTeam(event.Player.OppositeTeam())
		s.Stats(event.Player).BerriesKickedTheirTeam++
	}
}

func (event *CarryFoodEvent) Apply(s *StateMachine, time time.Time) {
	s.player(event.Player).HasBerry = true
}
