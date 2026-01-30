package models

// A model for all events that the KQ websocket sends.
// https://kqhivemind.com/wiki/Stats_Socket_Events is helpful documentation

type BerryDepositEvent struct {
	X      int
	Y      int
	Player PlayerId
}

type BerryKickInEvent struct {
	X           int
	Y           int
	Player      PlayerId
	PlayersHive bool
}

type BlessMaidenEvent struct {
	X    int
	Y    int
	Team TeamColor1
}

type CarryFoodEvent struct {
	Player PlayerId
}

type GameEndEvent struct {
	MapName     string
	GoldOnLeft  bool
	Duration    float64
	AttractMode bool
}

type GameStartEvent struct {
	MapName         string
	GoldOnLeft      bool
	ElapsedGameTime float64 // Always zero supposedly
	AttractMode     bool
	CabVersion      *string // Optional on older cabs
}

type GetOffSnailEvent struct {
	X      int
	Y      int
	Killer *PlayerId // nil if not killed
	Drone  PlayerId
}

type GetOnSnailEvent struct {
	X     int
	Y     int
	Drone PlayerId
}

type GlanceEvent struct {
	X       int
	Y       int
	Player1 PlayerId
	Player2 PlayerId
}

type MapStartEvent struct {
	MapName         string
	GoldOnLeft      bool
	ElapsedGameTime float64 // Always zero supposedly
	AttractMode     bool
	CabVersion      *string // Optional on older cabs
}

type PlayerKillEvent struct {
	X           int
	Y           int
	Killer      PlayerId
	Victim      PlayerId
	VictimsType VictimType
}

type PlayerNamesEvent struct {
	// There are 1-10 player names.  Just not a meaningful event
}

type ReserveMaidenEvent struct {
	X      int
	Y      int
	Player PlayerId
}

type SnailEatEvent struct {
	X      int
	Y      int
	Rider  PlayerId
	Victim PlayerId
}

type SnailEscapeEvent struct {
	X       int
	Y       int
	Escapee PlayerId
}

type SpawnEvent struct {
	Player PlayerId
	IsBot  bool
}

type UnreserveMaidenEvent struct {
	X      int
	Y      int
	Killer *PlayerId // Not always killed
	Drone  PlayerId
}

type UseMaidenEvent struct {
	X        int
	Y        int
	GateType GateType
	Player   PlayerId
}

type VictoryEvent struct {
	Team         TeamColor2
	WinCondition WinCondition
}
