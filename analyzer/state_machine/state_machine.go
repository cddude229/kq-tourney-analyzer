package state_machine

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
	"log"
)

type StateMachine struct {
	playerState map[models.PlayerId]*PlayerState
	playerStats map[models.PlayerId]*PlayerStats

	blueBerries      int
	goldBerries      int
	remainingBerries int

	winningTeam  *models.TeamColor2
	winCondition *models.WinCondition
}

type PlayerState struct {
	HasBerry   bool
	OnSnail    bool
	BeingEaten bool

	HasSpeed  bool
	IsWarrior bool

	IsBot bool
}

func (s *PlayerState) respawn() {
	s.HasBerry = false
	s.OnSnail = false
	s.BeingEaten = false

	s.HasSpeed = false
	s.IsWarrior = false

	// Don't touch IsBot
}

type PlayerStats struct {
	// Berries
	BerriesDunked          int
	BerriesKickedOurTeam   int
	BerriesKickedTheirTeam int

	// Gate usage
	GateDenyKills int
	KilledInGate  int
	LeftGate      int
}

func New() *StateMachine {
	return &StateMachine{
		playerState: make(map[models.PlayerId]*PlayerState),
		playerStats: make(map[models.PlayerId]*PlayerStats),

		blueBerries: 0,
		goldBerries: 0,

		remainingBerries: 0,
	}
}

func (s *StateMachine) player(playerId models.PlayerId) *PlayerState {
	playerState, ok := s.playerState[playerId]
	if !ok {
		playerState = &PlayerState{}
		playerState.respawn()
		s.playerState[playerId] = playerState
	}
	return playerState
}

func (s *StateMachine) stats(playerId models.PlayerId) *PlayerStats {
	stats, ok := s.playerStats[playerId]
	if !ok {
		stats = &PlayerStats{}
		s.playerStats[playerId] = stats
	}
	return stats
}

func (s *StateMachine) countBerry(color models.TeamColor2) {
	if color == models.BlueTeam2 {
		s.blueBerries++
	} else {
		s.goldBerries++
	}
}

// HandleHivemindEvent returns false if skipped or error, and true on success
func (s *StateMachine) HandleHivemindEvent(event hivemind.HivemindEvent) (bool, error) {
	if event.IsBerryDeposit() {
		e, err := event.BerryDeposit()
		if err != nil {
			return false, err
		}
		s.BerryDeposit(e)
	} else if event.IsBerryKickIn() {
		e, err := event.BerryKickIn()
		if err != nil {
			return false, err
		}
		s.BerryKickIn(e)
	} else if event.IsBlessMaiden() {
		e, err := event.BlessMaiden()
		if err != nil {
			return false, err
		}
		s.BlessMaiden(e)
	} else if event.IsCarryFood() {
		e, err := event.CarryFood()
		if err != nil {
			return false, err
		}
		s.CarryFood(e)
	} else if event.IsGameEnd() {
		e, err := event.GameEnd()
		if err != nil {
			return false, err
		}
		s.GameEnd(e)
	} else if event.IsGameStart() {
		e, err := event.GameStart()
		if err != nil {
			return false, err
		}
		s.GameStart(e)
	} else if event.IsGetOffSnail() {
		e, err := event.GetOffSnail()
		if err != nil {
			return false, err
		}
		s.GetOffSnail(e)
	} else if event.IsGetOnSnail() {
		e, err := event.GetOnSnail()
		if err != nil {
			return false, err
		}
		s.GetOnSnail(e)
	} else if event.IsGlance() {
		e, err := event.Glance()
		if err != nil {
			return false, err
		}
		s.Glance(e)
	} else if event.IsMapStart() {
		e, err := event.MapStart()
		if err != nil {
			return false, err
		}
		s.MapStart(e)
	} else if event.IsPlayerKill() {
		e, err := event.PlayerKill()
		if err != nil {
			return false, err
		}
		s.PlayerKill(e)
	} else if event.IsPlayerNames() {
		// Not handling this one
	} else if event.IsReserveMaiden() {
		e, err := event.ReserveMaiden()
		if err != nil {
			return false, err
		}
		s.ReserveMaiden(e)
	} else if event.IsSnailEat() {
		e, err := event.SnailEat()
		if err != nil {
			return false, err
		}
		s.SnailEat(e)
	} else if event.IsSnailEscape() {
		e, err := event.SnailEscape()
		if err != nil {
			return false, err
		}
		s.SnailEscape(e)
	} else if event.IsSpawn() {
		e, err := event.Spawn()
		if err != nil {
			return false, err
		}
		s.Spawn(e)
	} else if event.IsUnreserveMaiden() {
		e, err := event.UnreserveMaiden()
		if err != nil {
			return false, err
		}
		s.UnreserveMaiden(e)
	} else if event.IsUseMaiden() {
		e, err := event.UseMaiden()
		if err != nil {
			return false, err
		}
		s.UseMaiden(e)
	} else if event.IsVictory() {
		e, err := event.Victory()
		if err != nil {
			return false, err
		}
		s.Victory(e)
	} else if event.EventType == "cabinetOnline" {
		// Not sure what this one is actually
	} else {
		log.Printf("Unknown event: %s", event.EventType)
		return false, nil
	}

	return true, nil
}

func (s *StateMachine) BerryDeposit(event *models.BerryDepositEvent) {
	s.remainingBerries--

	s.player(event.Player).HasBerry = false

	s.countBerry(event.Player.Team())
	s.stats(event.Player).BerriesDunked++
}

func (s *StateMachine) BerryKickIn(event *models.BerryKickInEvent) {
	s.remainingBerries--

	if event.PlayersHive {
		s.countBerry(event.Player.Team())
		s.stats(event.Player).BerriesKickedOurTeam++
	} else {
		s.countBerry(event.Player.OppositeTeam())
		s.stats(event.Player).BerriesKickedTheirTeam++
	}
}

func (s *StateMachine) BlessMaiden(event *models.BlessMaidenEvent) {
	// TODO
}

func (s *StateMachine) CarryFood(event *models.CarryFoodEvent) {
	s.player(event.Player).HasBerry = true
}

func (s *StateMachine) GameEnd(event *models.GameEndEvent) {
	// TODO
}

func (s *StateMachine) GameStart(event *models.GameStartEvent) {
	// TODO
}

func (s *StateMachine) GetOffSnail(event *models.GetOffSnailEvent) {
	s.player(event.Drone).OnSnail = false

	// TODO: Record snail pixels
}

func (s *StateMachine) GetOnSnail(event *models.GetOnSnailEvent) {
	s.player(event.Drone).OnSnail = true

	// TODO
}

func (s *StateMachine) Glance(event *models.GlanceEvent) {
	// TODO
}

func (s *StateMachine) MapStart(event *models.MapStartEvent) {
	// TODO
}

func (s *StateMachine) PlayerKill(event *models.PlayerKillEvent) {
	// TODO
}

func (s *StateMachine) ReserveMaiden(event *models.ReserveMaidenEvent) {
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

func (s *StateMachine) Spawn(event *models.SpawnEvent) {
	player := s.player(event.Player)
	player.respawn()
	player.IsBot = event.IsBot
}

func (s *StateMachine) UnreserveMaiden(event *models.UnreserveMaidenEvent) {
	if event.Killer != nil {
		s.stats(*event.Killer).GateDenyKills++
		s.stats(event.Drone).KilledInGate++
	} else {
		// TODO: Might have been bumped out?  Compare with glance events
		s.stats(event.Drone).LeftGate++
	}
}

func (s *StateMachine) UseMaiden(event *models.UseMaidenEvent) {
	s.remainingBerries--

	if event.GateType == models.SpeedGate {
		s.player(event.Player).HasSpeed = true
	} else {
		s.player(event.Player).IsWarrior = true
	}
}

func (s *StateMachine) Victory(event *models.VictoryEvent) {
	s.winningTeam = &event.Team
	s.winCondition = &event.WinCondition
}
