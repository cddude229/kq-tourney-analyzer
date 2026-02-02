package models

import "time"

func (event *GlanceEvent) Apply(s *StateMachine, time time.Time) {
	p1 := s.player(event.Player1).stateArrayIdx()
	p2 := s.player(event.Player2).stateArrayIdx()

	s.Stats(event.Player1).BumpCounter[p1][p2]++
	s.Stats(event.Player2).BumpCounter[p2][p1]++
}

func (event *PlayerKillEvent) Apply(s *StateMachine, time time.Time) {
	killerStateIdx := s.player(event.Killer).stateArrayIdx()
	victimState := s.player(event.Victim)
	victimStats := s.Stats(event.Victim)
	victimStateIdx := victimState.stateArrayIdx()

	s.Stats(event.Killer).KillCounter[killerStateIdx][victimStateIdx]++
	victimStats.DeathCounter[killerStateIdx][victimStateIdx]++

	if victimState.IsWarrior {
		if victimState.IsSpeed {
			victimStats.SpeedWarriorUptime += time.UnixMilli() - victimState.BecameWarriorAt.UnixMilli()
		} else {
			victimStats.VanillaWarriorUptime += time.UnixMilli() - victimState.BecameWarriorAt.UnixMilli()
		}
	} else if victimState.IsSpeed {
		victimStats.SpeedDroneUptime += time.UnixMilli() - victimState.GotSpeedAt.UnixMilli()
	}

	victimState.respawn()
}

func (event *SpawnEvent) Apply(s *StateMachine, time time.Time) {
	player := s.player(event.Player)
	player.respawn()
	player.IsBot = event.IsBot
}
