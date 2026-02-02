package models

import (
	"time"
)

type GateStateAndStats struct {
	Color      *TeamColor1
	LastTagged time.Time

	gateType      *GateType
	gateTypeGuess GateType // Always set during creation

	TimesUsed int

	TimeForBlue int64
	TimeForGold int64
}

func (s *GateStateAndStats) IsWarriorGate() bool {
	if s.gateType == nil {
		return s.gateTypeGuess == WarriorGate
	}
	return *s.gateType == WarriorGate
}
func (s *GateStateAndStats) IsSpeedGate() bool {
	if s.gateType == nil {
		return s.gateTypeGuess == SpeedGate
	}
	return *s.gateType == SpeedGate
}

func (s *StateMachine) Gate(x int, y int) *GateStateAndStats {
	key := x * y
	gate, exists := s.gates[key]
	if !exists {
		gate = &GateStateAndStats{
			// TODO: Use x+y and the map to actually guess here
			gateTypeGuess: WarriorGate,
		}
		s.gates[key] = gate
	}
	return gate
}

// CalculateGateControlTimeInSeconds returns (Blue, Gold)
func (s *StateMachine) CalculateGateControlTimeInSeconds() (float64, float64) {
	var timeBlueTotal, timeGoldTotal int64 = 0, 0

	for _, gate := range s.gates {
		timeBlueTotal += gate.TimeForBlue
		timeGoldTotal += gate.TimeForGold
	}

	return float64(timeBlueTotal) / 1000.0, float64(timeGoldTotal) / 1000.0
}

// CalculateWarriorGateControlTimeInSeconds returns (Blue, Gold)
func (s *StateMachine) CalculateWarriorGateControlTimeInSeconds() (float64, float64) {
	var timeBlueTotal, timeGoldTotal int64 = 0, 0

	for _, gate := range s.gates {
		if gate.IsWarriorGate() {
			timeBlueTotal += gate.TimeForBlue
			timeGoldTotal += gate.TimeForGold
		}
	}

	return float64(timeBlueTotal) / 1000.0, float64(timeGoldTotal) / 1000.0
}

// CalculateSpeedGateControlTimeInSeconds returns (Blue, Gold)
func (s *StateMachine) CalculateSpeedGateControlTimeInSeconds() (float64, float64) {
	var timeBlueTotal, timeGoldTotal int64 = 0, 0

	for _, gate := range s.gates {
		if gate.IsSpeedGate() {
			timeBlueTotal += gate.TimeForBlue
			timeGoldTotal += gate.TimeForGold
		}
	}

	return float64(timeBlueTotal) / 1000.0, float64(timeGoldTotal) / 1000.0
}
