package models

import "time"

func (event *GlanceEvent) Apply(s *StateMachine, time time.Time) {
	p1 := s.player(event.Player1).stateArrayIdx()
	p2 := s.player(event.Player2).stateArrayIdx()

	s.stats(event.Player1).BumpCounter[p1][p2]++
	s.stats(event.Player2).BumpCounter[p2][p1]++
}

func (event *PlayerKillEvent) Apply(s *StateMachine, time time.Time) {
	killerStateIdx := s.player(event.Killer).stateArrayIdx()
	victimState := s.player(event.Victim)
	victimStateIdx := victimState.stateArrayIdx()

	s.stats(event.Killer).KillCounter[killerStateIdx][victimStateIdx]++
	s.stats(event.Victim).DeathCounter[killerStateIdx][victimStateIdx]++

	victimState.respawn()
}

func (event *SpawnEvent) Apply(s *StateMachine, time time.Time) {
	player := s.player(event.Player)
	player.respawn()
	player.IsBot = event.IsBot
}
