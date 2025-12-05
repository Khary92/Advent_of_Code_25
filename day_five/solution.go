package day_five

import (
	"AoC25/helper"
	"fmt"
	"strconv"
	"strings"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

type Range struct {
	Lower int
	Upper int
}

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

func SolveThaThingPart2() {
	_, ranges := SeparateTheThing()

	thaFreshCount := Set[int]{}
	for _, rangeValues := range ranges {
		fmt.Printf("Tha fresh count is %d\n", thaFreshCount)
		for i := rangeValues.Lower; i <= rangeValues.Upper; i++ {
			thaFreshCount.Add(i)
		}
	}

	fmt.Printf("Tha fresh count is %d\n", len(thaFreshCount))
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

			ranges = append(ranges, Range{lower, upper})
		}

		checkableValue, _ := strconv.Atoi(line)
		values = append(values, checkableValue)

	}
	return
}
