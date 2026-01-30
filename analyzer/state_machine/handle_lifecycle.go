package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) GameEnd(event *models.GameEndEvent) {
	// TODO
}

func (s *StateMachine) GameStart(event *models.GameStartEvent) {
	// TODO
}

func (s *StateMachine) MapStart(event *models.MapStartEvent) {
	// TODO
}

func (s *StateMachine) Victory(event *models.VictoryEvent) {
	s.winningTeam = &event.Team
	s.winCondition = &event.WinCondition
}
