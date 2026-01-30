package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) Glance(event *models.GlanceEvent) {
	// TODO
}

func (s *StateMachine) PlayerKill(event *models.PlayerKillEvent) {
	// TODO
}

func (s *StateMachine) Spawn(event *models.SpawnEvent) {
	player := s.player(event.Player)
	player.respawn()
	player.IsBot = event.IsBot
}
