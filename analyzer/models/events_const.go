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

// TeamColor2 is used exclusively in VictoryEvent
type TeamColor2 string

const BlueTeam2 TeamColor2 = "Blue"
const GoldTeam2 TeamColor2 = "Gold"

type WinCondition string

const EconomicVictory WinCondition = "Economic"
const MilitaryVictory WinCondition = "Military"
const SnailVictory WinCondition = "Snail"
