package hivemind

import (
	"cddude229/kq-tourney-analyzer/models"
	"strconv"
	"strings"
)

func parseXY(values []string, xField int, yField int) (int, int, error) {
	x, err := strconv.Atoi(values[xField])
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(values[yField])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func parsePlayer(values []string, field int) (models.PlayerId, error) {
	id, err := strconv.Atoi(values[field])
	if err != nil {
		return models.PlayerId(0), err
	}

	return models.PlayerId(id), nil
}

func parseOptionalPlayer(values []string, field int) (*models.PlayerId, error) {
	if values[field] == "\"\"" {
		return nil, nil
	}

	id, err := strconv.Atoi(values[field])
	if err != nil {
		return nil, err
	}

	pid := models.PlayerId(id)
	return &pid, nil
}

func parseBool(values []string, field int) (bool, error) {
	b, err := strconv.ParseBool(values[field])
	return b, err
}

func parseFloat(values []string, field int) (float64, error) {
	f, err := strconv.ParseFloat(values[field], 64)
	return f, err
}

func (e *HivemindEvent) parseValues() []string {
	return strings.Split(strings.Trim(e.Values, "{}"), ",")
}

func (e *HivemindEvent) IsBerryDeposit() bool {
	return e.EventType == "berryDeposit"
}
func (e *HivemindEvent) BerryDeposit() (*models.BerryDepositEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	player, err := parsePlayer(values, 2)
	if err != nil {
		return nil, err
	}

	return &models.BerryDepositEvent{
		X:      x,
		Y:      y,
		Player: player,
	}, nil
}

func (e *HivemindEvent) IsBerryKickIn() bool {
	return e.EventType == "berryKickIn"
}
func (e *HivemindEvent) BerryKickIn() (*models.BerryKickInEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	player, err := parsePlayer(values, 2)
	if err != nil {
		return nil, err
	}

	playersHive, err := parseBool(values, 3)
	if err != nil {
		return nil, err
	}

	return &models.BerryKickInEvent{
		X:           x,
		Y:           y,
		Player:      player,
		PlayersHive: playersHive,
	}, nil
}

func (e *HivemindEvent) IsBlessMaiden() bool {
	return e.EventType == "blessMaiden"
}
func (e *HivemindEvent) BlessMaiden() (*models.BlessMaidenEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	return &models.BlessMaidenEvent{
		X:    x,
		Y:    y,
		Team: models.TeamColor1(values[2]),
	}, nil
}

func (e *HivemindEvent) IsCarryFood() bool {
	return e.EventType == "carryFood"
}
func (e *HivemindEvent) CarryFood() (*models.CarryFoodEvent, error) {
	values := e.parseValues()

	player, err := parsePlayer(values, 0)
	if err != nil {
		return nil, err
	}

	return &models.CarryFoodEvent{
		Player: player,
	}, nil
}

func (e *HivemindEvent) IsGameEnd() bool {
	return e.EventType == "gameend"
}
func (e *HivemindEvent) GameEnd() (*models.GameEndEvent, error) {
	values := e.parseValues()

	unknown1, err := parseBool(values, 1)
	if err != nil {
		return nil, err
	}

	duration, err := parseFloat(values, 2)
	if err != nil {
		return nil, err
	}

	unknown2, err := parseBool(values, 3)
	if err != nil {
		return nil, err
	}

	return &models.GameEndEvent{
		MapName:  values[0],
		Unknown1: unknown1,
		Duration: duration,
		Unknown2: unknown2,
	}, nil
}

func (e *HivemindEvent) IsGameStart() bool {
	return e.EventType == "gamestart"
}
func (e *HivemindEvent) GameStart() (*models.GameStartEvent, error) {
	values := e.parseValues()

	goldOnLeft, err := parseBool(values, 1)
	if err != nil {
		return nil, err
	}

	duration, err := parseFloat(values, 2)
	if err != nil {
		return nil, err
	}

	attractMode, err := parseBool(values, 3)
	if err != nil {
		return nil, err
	}

	var cabVersion *string
	if len(values) > 4 {
		cabVersion = &values[4]
	}

	return &models.GameStartEvent{
		MapName:         values[0],
		GoldOnLeft:      goldOnLeft,
		ElapsedGameTime: duration,
		AttractMode:     attractMode,
		CabVersion:      cabVersion,
	}, nil
}

func (e *HivemindEvent) IsGetOffSnail() bool {
	return e.EventType == "getOffSnail"
}
func (e *HivemindEvent) GetOffSnail() (*models.GetOffSnailEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	killer, err := parseOptionalPlayer(values, 2)
	if err != nil {
		return nil, err
	}

	drone, err := parsePlayer(values, 3)
	if err != nil {
		return nil, err
	}

	return &models.GetOffSnailEvent{
		X:      x,
		Y:      y,
		Killer: killer,
		Drone:  drone,
	}, nil
}

func (e *HivemindEvent) IsGetOnSnail() bool {
	return e.EventType == "getOnSnail"
}
func (e *HivemindEvent) GetOnSnail() (*models.GetOnSnailEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	player, err := parsePlayer(values, 2)
	if err != nil {
		return nil, err
	}

	return &models.GetOnSnailEvent{
		X:     x,
		Y:     y,
		Drone: player,
	}, nil
}

func (e *HivemindEvent) IsGlance() bool {
	return e.EventType == "glance"
}
func (e *HivemindEvent) Glance() (*models.GlanceEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	player1, err := parsePlayer(values, 2)
	if err != nil {
		return nil, err
	}

	player2, err := parsePlayer(values, 3)
	if err != nil {
		return nil, err
	}

	return &models.GlanceEvent{
		X:       x,
		Y:       y,
		Player1: player1,
		Player2: player2,
	}, nil
}

func (e *HivemindEvent) IsMapStart() bool {
	return e.EventType == "mapstart"
}
func (e *HivemindEvent) MapStart() (*models.MapStartEvent, error) {
	values := e.parseValues()

	goldOnLeft, err := parseBool(values, 1)
	if err != nil {
		return nil, err
	}

	duration, err := parseFloat(values, 2)
	if err != nil {
		return nil, err
	}

	attractMode, err := parseBool(values, 3)
	if err != nil {
		return nil, err
	}

	var cabVersion *string
	if len(values) > 4 {
		cabVersion = &values[4]
	}

	return &models.MapStartEvent{
		MapName:         values[0],
		GoldOnLeft:      goldOnLeft,
		ElapsedGameTime: duration,
		AttractMode:     attractMode,
		CabVersion:      cabVersion,
	}, nil
}

func (e *HivemindEvent) IsPlayerKill() bool {
	return e.EventType == "playerKill"
}
func (e *HivemindEvent) PlayerKill() (*models.PlayerKillEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	killer, err := parsePlayer(values, 2)
	if err != nil {
		return nil, err
	}

	victim, err := parsePlayer(values, 3)
	if err != nil {
		return nil, err
	}

	return &models.PlayerKillEvent{
		X:           x,
		Y:           y,
		Killer:      killer,
		Victim:      victim,
		VictimsType: models.VictimType(values[4]),
	}, nil
}

func (e *HivemindEvent) IsPlayerNames() bool {
	return e.EventType == "playernames"
}
func (e *HivemindEvent) PlayerNames() (*models.PlayerNamesEvent, error) {
	// Not actually implementing this
	return &models.PlayerNamesEvent{}, nil
}

func (e *HivemindEvent) IsReserveMaiden() bool {
	return e.EventType == "reserveMaiden"
}
func (e *HivemindEvent) ReserveMaiden() (*models.ReserveMaidenEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	player, err := parsePlayer(values, 2)
	if err != nil {
		return nil, err
	}

	return &models.ReserveMaidenEvent{
		X:      x,
		Y:      y,
		Player: player,
	}, nil
}

func (e *HivemindEvent) IsSnailEat() bool {
	return e.EventType == "snailEat"
}
func (e *HivemindEvent) SnailEat() (*models.SnailEatEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	rider, err := parsePlayer(values, 2)
	if err != nil {
		return nil, err
	}

	victim, err := parsePlayer(values, 3)
	if err != nil {
		return nil, err
	}

	return &models.SnailEatEvent{
		X:      x,
		Y:      y,
		Rider:  rider,
		Victim: victim,
	}, nil
}

func (e *HivemindEvent) IsSnailEscape() bool {
	return e.EventType == "snailEscape"
}
func (e *HivemindEvent) SnailEscape() (*models.SnailEscapeEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	escapee, err := parsePlayer(values, 2)
	if err != nil {
		return nil, err
	}

	return &models.SnailEscapeEvent{
		X:       x,
		Y:       y,
		Escapee: escapee,
	}, nil
}

func (e *HivemindEvent) IsSpawn() bool {
	return e.EventType == "spawn"
}
func (e *HivemindEvent) Spawn() (*models.SpawnEvent, error) {
	values := e.parseValues()

	player, err := parsePlayer(values, 0)
	if err != nil {
		return nil, err
	}

	isBot, err := parseBool(values, 1)
	if err != nil {
		return nil, err
	}

	return &models.SpawnEvent{
		Player: player,
		IsBot:  isBot,
	}, nil
}

func (e *HivemindEvent) IsUnreserveMaiden() bool {
	return e.EventType == "unreserveMaiden"
}
func (e *HivemindEvent) UnreserveMaiden() (*models.UnreserveMaidenEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	killer, err := parseOptionalPlayer(values, 2)
	if err != nil {
		return nil, err
	}

	drone, err := parsePlayer(values, 3)
	if err != nil {
		return nil, err
	}

	return &models.UnreserveMaidenEvent{
		X:      x,
		Y:      y,
		Killer: killer,
		Drone:  drone,
	}, nil
}

func (e *HivemindEvent) IsUseMaiden() bool {
	return e.EventType == "useMaiden"
}
func (e *HivemindEvent) UseMaiden() (*models.UseMaidenEvent, error) {
	values := e.parseValues()

	x, y, err := parseXY(values, 0, 1)
	if err != nil {
		return nil, err
	}

	player, err := parsePlayer(values, 3)
	if err != nil {
		return nil, err
	}

	return &models.UseMaidenEvent{
		X:        x,
		Y:        y,
		GateType: models.GateType(values[2]),
		Player:   player,
	}, nil
}

func (e *HivemindEvent) IsVictory() bool {
	return e.EventType == "victory"
}
func (e *HivemindEvent) Victory() (*models.VictoryEvent, error) {
	values := e.parseValues()
	return &models.VictoryEvent{
		Team:         models.TeamColor2(values[0]),
		WinCondition: models.WinCondition(values[1]),
	}, nil
}
