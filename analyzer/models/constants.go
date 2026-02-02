package models

type PlayerId int

const (
	unused PlayerId = iota
	GoldQueen
	BlueQueen
	GoldStripes
	BlueStripes
	GoldAbs
	BlueAbs
	GoldSkulls
	BlueSkulls
	GoldChex
	BlueChex
)

func (p PlayerId) Team() TeamColor2 {
	if p%2 == 0 {
		return BlueTeam2
	}

	return GoldTeam2
}

func (p PlayerId) IsBlue() bool {
	return p.Team() == BlueTeam2
}

func (p PlayerId) OppositeTeam() TeamColor2 {
	if p%2 == 1 {
		return BlueTeam2
	}

	return GoldTeam2
}

func (p PlayerId) IsQueen() bool {
	return p == GoldQueen || p == BlueQueen
}

// TeamColor1 is used exclusively in BlessMaidenEvent
type TeamColor1 string

const BlueTeam1 TeamColor1 = "Blue"
const GoldTeam1 TeamColor1 = "Red"

type VictimType string

const Queen VictimType = "Queen"
const Warrior VictimType = "Soldier"
const Drone VictimType = "Worker"

type GateType string

const SpeedGate GateType = "maiden_speed"
const WarriorGate GateType = "maiden_wings"

// TeamColor2 is used exclusively in VictoryEvent, but reused elsewhere
type TeamColor2 string

const BlueTeam2 TeamColor2 = "Blue"
const GoldTeam2 TeamColor2 = "Gold"

type WinCondition string

const EconomicVictory WinCondition = "Economic"
const MilitaryVictory WinCondition = "Military"
const SnailVictory WinCondition = "Snail"
