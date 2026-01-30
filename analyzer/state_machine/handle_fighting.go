package state_machine

import "cddude229/kq-tourney-analyzer/models"

func (s *StateMachine) Glance(event *models.GlanceEvent) {
	p1 := s.player(event.Player1)
	p2 := s.player(event.Player2)

	s.stats(event.Player1).BumpCounter[p1.stateArrayIdx()][p2.stateArrayIdx()]++
	s.stats(event.Player2).BumpCounter[p2.stateArrayIdx()][p1.stateArrayIdx()]++
}

func (s *StateMachine) PlayerKill(event *models.PlayerKillEvent) {
	killerStateIdx := s.player(event.Killer).stateArrayIdx()
	victimState := s.player(event.Victim)
	victimStateIdx := victimState.stateArrayIdx()

	s.stats(event.Killer).KillCounter[killerStateIdx][victimStateIdx]++
	s.stats(event.Victim).DeathCounter[killerStateIdx][victimStateIdx]++

	victimState.respawn()
}

func (s *StateMachine) Spawn(event *models.SpawnEvent) {
	player := s.player(event.Player)
	player.respawn()
	player.IsBot = event.IsBot
}
