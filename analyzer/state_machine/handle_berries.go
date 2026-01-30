package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) countBerryForTeam(color models.TeamColor2) {
	if color == models.BlueTeam2 {
		s.blueBerries++
	} else {
		s.goldBerries++
	}
}

func (s *StateMachine) BerryDeposit(event *models.BerryDepositEvent) {
	s.remainingBerries--

	s.player(event.Player).HasBerry = false

	s.countBerryForTeam(event.Player.Team())
	s.stats(event.Player).BerriesDunked++
}

func (s *StateMachine) BerryKickIn(event *models.BerryKickInEvent) {
	s.remainingBerries--

	if event.PlayersHive {
		s.countBerryForTeam(event.Player.Team())
		s.stats(event.Player).BerriesKickedOurTeam++
	} else {
		s.countBerryForTeam(event.Player.OppositeTeam())
		s.stats(event.Player).BerriesKickedTheirTeam++
	}
}

func (s *StateMachine) CarryFood(event *models.CarryFoodEvent) {
	s.player(event.Player).HasBerry = true
}
