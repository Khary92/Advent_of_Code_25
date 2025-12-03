package day_three

import (
	"AoC25/helper"
	"fmt"
)

func SolveThaThing() {
	lines, err := helper.ReadLinesAsIntArrays("day_three/input.txt")

	if err != nil {
		return
	}

	count := 0
	for _, line := range lines {
		result := CalculateThaThing(line)
		count += result.GetThaJoltage()
	}

	fmt.Println(count)
}
