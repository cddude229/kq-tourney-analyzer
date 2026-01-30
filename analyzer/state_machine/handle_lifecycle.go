package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) GameEnd(event *models.GameEndEvent) {
	s.mapName = event.MapName
	s.finalGameDuration = event.Duration

	// TODO: Add estimated snail distance to the rider(s)
}

func (s *StateMachine) GameStart(event *models.GameStartEvent) {
	s.mapName = event.MapName
	s.goldOnLeft = event.GoldOnLeft
	s.cabVersion = event.CabVersion
}

func (s *StateMachine) MapStart(event *models.MapStartEvent) {
	s.mapName = event.MapName
	s.goldOnLeft = event.GoldOnLeft
	s.cabVersion = event.CabVersion
}

func (s *StateMachine) Victory(event *models.VictoryEvent) {
	s.winningTeam = &event.Team
	s.winCondition = &event.WinCondition
}
