package day_five

import (
	"AoC25/helper"
	"fmt"
	"strconv"
	"strings"
)

func SolveThaThing() {
	values, ranges := SeparateTheThing()

	thaFreshCount := 0
	for _, value := range values {
		for _, rangeValues := range ranges {
			if value >= rangeValues.Lower && value <= rangeValues.Upper {
				thaFreshCount++
				break
			}
		}
	}

	fmt.Printf("Tha fresh count is %d\n", thaFreshCount)
}

func SeparateTheThing() (values []int, ranges []Range) {
	lines, err := helper.ReadLines("day_five/input.txt")

	if err != nil {
		return
	}

	for _, line := range lines {
		fmt.Println(line)
		if line == "" {
			continue
		}

		//AddRanges
		if strings.ContainsRune(line, '-') {
			upperAndLower := helper.SplitByChars(line, "-")

			lower, _ := strconv.Atoi(upperAndLower[0])
			upper, _ := strconv.Atoi(upperAndLower[1])

			ranges = append(ranges, Range{lower, upper, 0})
		}

		checkableValue, _ := strconv.Atoi(line)
		values = append(values, checkableValue)

	}
	return
}
