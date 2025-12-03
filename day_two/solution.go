package day_two

import (
	"AoC25/helper"
	"fmt"
	"strconv"
)

func SolveThaThing() {
	input, err := helper.ReadLines("day_two/input.txt")

	if err != nil {
		fmt.Println(err)
		return

	}

	splitStrings := helper.SplitByChars(input[0], ",")

	var foundInvalidValuesList []int

	for _, values := range splitStrings {

		tuple := helper.SplitByChars(values, "-")
		foundInvalidValuesList = append(foundInvalidValuesList, IterateRange(tuple)...)
	}

	result := 0
	for i := 0; i < len(foundInvalidValuesList); i++ {
		result += foundInvalidValuesList[i]
	}

	fmt.Println("Result: " + strconv.Itoa(result))
}

func IterateRange(input []string) []int {
	var results []int
	lower, lowerErr := strconv.Atoi(input[0])
	upper, upperErr := strconv.Atoi(input[1])

	if lowerErr != nil || upperErr != nil {
		fmt.Println("Some text was not a valid int i guess.")
		return nil
	}

	// Iterate range of tuples
	for i := lower; i < upper+1; i++ {
		isDuplicate := ValidateStringFormat(strconv.Itoa(i))
		if isDuplicate {
			fmt.Println("Found a duplicate entry: " + strconv.Itoa(i))
			results = append(results, i)
		}
	}

	return results
}

func ValidateStringFormat(value string) bool {
	if len(value)%2 != 0 {
		return false
	}
	splitIndex := len(value) / 2

	first := value[:splitIndex]
	second := value[splitIndex:]

	return first == second
}
