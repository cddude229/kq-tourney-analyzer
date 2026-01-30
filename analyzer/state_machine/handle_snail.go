package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) GetOffSnail(event *models.GetOffSnailEvent) {
	s.player(event.Drone).OnSnail = false

	// TODO: Record snail pixels
}

func (s *StateMachine) GetOnSnail(event *models.GetOnSnailEvent) {
	s.player(event.Drone).OnSnail = true

	// TODO
}

func (s *StateMachine) SnailEat(event *models.SnailEatEvent) {
	s.player(event.Victim).BeingEaten = true
	// TODO: Update the snail position
	// TODO: Sac stats
}

func (s *StateMachine) SnailEscape(event *models.SnailEscapeEvent) {
	s.player(event.Escapee).BeingEaten = false
	// TODO: Sac stats
}
