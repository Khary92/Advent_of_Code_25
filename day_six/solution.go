package day_six

import (
	"AoC25/helper"
	"fmt"
)

func SolveThaThing() {
	lines, err := helper.ReadLines("day_six/input.txt")

	if err != nil {
		return
	}

	resultHodlers := DoTheRightStructure(lines)

	result := 0

	for _, hodler := range resultHodlers {
		hodler.DoTheMath()
		result += hodler.Result
	}

	fmt.Println(result)
}
