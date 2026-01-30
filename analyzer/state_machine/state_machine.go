package state_machine

import (
	"cddude229/kq-tourney-analyzer/hivemind"
	"cddude229/kq-tourney-analyzer/models"
	"log"
)

type StateMachine struct {
	// TODO
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
		e, err := event.PlayerNames()
		if err != nil {
			return false, err
		}
		s.PlayerNames(e)
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
	// TODO
}

func (s *StateMachine) BerryKickIn(event *models.BerryKickInEvent) {
	// TODO
}

func (s *StateMachine) BlessMaiden(event *models.BlessMaidenEvent) {
	// TODO
}

func (s *StateMachine) CarryFood(event *models.CarryFoodEvent) {
	// TODO
}

func (s *StateMachine) GameEnd(event *models.GameEndEvent) {
	// TODO
}

func (s *StateMachine) GameStart(event *models.GameStartEvent) {
	// TODO
}

func (s *StateMachine) GetOffSnail(event *models.GetOffSnailEvent) {
	// TODO
}

func (s *StateMachine) GetOnSnail(event *models.GetOnSnailEvent) {
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

func (s *StateMachine) PlayerNames(event *models.PlayerNamesEvent) {
	// noop
}

func (s *StateMachine) ReserveMaiden(event *models.ReserveMaidenEvent) {
	// TODO
}

func (s *StateMachine) SnailEat(event *models.SnailEatEvent) {
	// TODO
}

func (s *StateMachine) SnailEscape(event *models.SnailEscapeEvent) {
	// TODO
}

func (s *StateMachine) Spawn(event *models.SpawnEvent) {
	// TODO
}

func (s *StateMachine) UnreserveMaiden(event *models.UnreserveMaidenEvent) {
	// TODO
}

func (s *StateMachine) UseMaiden(event *models.UseMaidenEvent) {
	// TODO
}

func (s *StateMachine) Victory(event *models.VictoryEvent) {
	// TODO
}
