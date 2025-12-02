package day_two

import (
	"AoC25/helper"
	"fmt"
	"strconv"
	"strings"
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
	value = strings.TrimSpace(value)
	chars := []rune(value)

	//if single digit it is false
	if len(chars) < 2 {
		return false
	}

	//if it is two digits the first needs to be the same as the second
	if len(chars) == 2 {
		return chars[0] == chars[1]
	}

	//if there are more that two it is possible that they are all the same. If yes its true
	allSame := true
	for i := 1; i < len(chars); i++ {
		if chars[i] != chars[0] {
			allSame = false
			break
		}
	}

	if allSame {
		return true
	}

	//check different batch sizes for pattern repetition
	for size := 1; size <= len(chars)/2; size++ {
		// if they cant be split in even chunks there will be no equality
		if len(chars)%size != 0 {
			continue
		}

		isValid := true
		for offset := size; offset < len(chars); offset += size {
			for j := 0; j < size; j++ {
				if chars[j] != chars[offset+j] {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}

		if isValid {
			return true
		}
	}

	return false
}
