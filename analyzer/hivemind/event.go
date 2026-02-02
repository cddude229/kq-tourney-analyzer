package hivemind

import (
	"cddude229/kq-tourney-analyzer/models"
	"fmt"
	"time"
)

type HivemindEvent struct {
	Id        int64
	Timestamp time.Time
	EventType string
	Values    string
	GameId    int64
}

type HivemindEventById []HivemindEvent

// TODO: Sorting UTs
func (a HivemindEventById) Len() int           { return len(a) }
func (a HivemindEventById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a HivemindEventById) Less(i, j int) bool { return a[i].Id < a[j].Id }

type SMEvent interface {
	Apply(s *models.StateMachine, time time.Time)
}

type noopEvent struct{}

func (e *noopEvent) Apply(s *models.StateMachine, time time.Time) {}

func (e *HivemindEvent) ToSMEvent() (SMEvent, error) {
	switch e.EventType {
	case "berryDeposit":
		return e.BerryDeposit()
	case "berryKickIn":
		return e.BerryKickIn()
	case "blessMaiden":
		return e.BlessMaiden()
	case "carryFood":
		return e.CarryFood()
	case "gameend":
		return e.GameEnd()
	case "gamestart":
		return e.GameStart()
	case "getOffSnail":
		return e.GetOffSnail()
	case "getOnSnail":
		return e.GetOnSnail()
	case "glance":
		return e.Glance()
	case "mapstart":
		return e.MapStart()
	case "playerKill":
		return e.PlayerKill()
	case "reserveMaiden":
		return e.ReserveMaiden()
	case "snailEat":
		return e.SnailEat()
	case "snailEscape":
		return e.SnailEscape()
	case "spawn":
		return e.Spawn()
	case "unreserveMaiden":
		return e.UnreserveMaiden()
	case "useMaiden":
		return e.UseMaiden()
	case "victory":
		return e.Victory()
	case "playernames", "cabinetOnline":
		return &noopEvent{}, nil
	default:
		return nil, fmt.Errorf("invalid event type: %s", e.EventType)
	}
}
