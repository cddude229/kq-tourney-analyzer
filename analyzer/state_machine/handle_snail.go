package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) GetOffSnail(event *models.GetOffSnailEvent) {
	rider := s.player(event.Drone)
	rider.OnSnail = false
	s.stats(event.Drone).recordSnailDistance(event.X - rider.LastRecordedSnailX)
	rider.LastRecordedSnailX = event.X
}

func (s *StateMachine) GetOnSnail(event *models.GetOnSnailEvent) {
	player := s.player(event.Drone)
	player.OnSnail = true
	player.LastRecordedSnailX = event.X
}

func (s *StateMachine) SnailEat(event *models.SnailEatEvent) {
	s.player(event.Victim).BeingEaten = true

	rider := s.player(event.Rider)
	s.stats(event.Rider).recordSnailDistance(event.X - rider.LastRecordedSnailX)
	rider.LastRecordedSnailX = event.X

	// TODO: Sac / eat stats
}

func (s *StateMachine) SnailEscape(event *models.SnailEscapeEvent) {
	s.player(event.Escapee).BeingEaten = false

	// TODO: Sac stats
}
