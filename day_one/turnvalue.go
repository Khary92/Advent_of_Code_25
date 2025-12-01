package day_one

import "strconv"

type TurnValue struct {
	Direction string
	Value     int
}

func ConvertToTurnValue(s string) TurnValue {
	turnValue := TurnValue{}

	chars := []rune(s)
	if chars[0] == 'R' {
		turnValue.Direction = "R"
	}

	if chars[0] == 'L' {
		turnValue.Direction = "L"
	}

	turnValue.Value, _ = strconv.Atoi(string(chars[1:]))

	return turnValue
}
