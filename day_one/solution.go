package day_one

import (
	"AoC25/helper"
	"fmt"
	"strconv"
)

func SolveThaThing() {
	lines, err := helper.ReadLines("day_one/input.txt")

	if err != nil {
		return
	}

	counter := 0
	dial := Dial{Value: 50}
	for _, line := range lines {
		token := ConvertToTurnValue(line)

		if token.Direction == "R" {
			dial.TurnRight(token.Value)
		}

		if token.Direction == "L" {
			dial.TurnLeft(token.Value)
		}

		if dial.Value == 0 {
			counter++
		}

		fmt.Println(counter)
	}
}

func SolveThaThing_Part2() {
	lines, err := helper.ReadLines("day_one/input.txt")

	if err != nil {
		return
	}

	counter := 0
	dial := Dial{Value: 50}
	for _, line := range lines {
		fmt.Println("Input: " + line)
		token := ConvertToTurnValue(line)

		if token.Direction == "R" {
			result := dial.TurnRight(token.Value)
			fmt.Println("Add " + strconv.Itoa(result))
			counter += result
		}

		if token.Direction == "L" {
			result := dial.TurnLeft(token.Value)
			fmt.Println("Add " + strconv.Itoa(result))
			counter += result
		}

		fmt.Println("Current value " + strconv.Itoa(dial.Value))
	}

	fmt.Println("Result " + strconv.Itoa(counter))
}
