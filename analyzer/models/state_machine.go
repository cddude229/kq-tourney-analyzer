package models

type StateMachine struct {
	playerState map[PlayerId]*PlayerState
	PlayerStats map[PlayerId]*PlayerStats

	gates map[int]*GateStateAndStats

	mapName     string
	goldOnLeft  bool
	attractMode bool
	cabVersion  *string

	blueBerries      int
	goldBerries      int
	remainingBerries int

	winningTeam       *TeamColor2
	winCondition      *WinCondition
	finalGameDuration float64 // uninitialized if not final value
}

func New() *StateMachine {
	return &StateMachine{
		playerState: make(map[PlayerId]*PlayerState),
		PlayerStats: make(map[PlayerId]*PlayerStats),

		gates: make(map[int]*GateStateAndStats),
	}
}
