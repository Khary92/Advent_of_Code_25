package day_three

import (
	"strconv"
)

type BankWithExquisiteBatteries struct {
	TheChosenTwelve []int
}

func RemoveLowestValuesFromFront(input []int) (result BankWithExquisiteBatteries) {
	targetCount := 12
	removeCount := len(input) - targetCount

	stack := make([]int, 0, len(input))

	for _, value := range input {
		for removeCount > 0 && len(stack) > 0 && stack[len(stack)-1] < value {
			stack = stack[:len(stack)-1]
			removeCount--
		}
		stack = append(stack, value)
	}

	if len(stack) > targetCount {
		stack = stack[:targetCount]
	}

	result.TheChosenTwelve = stack
	return
}

func (bank *BankWithExquisiteBatteries) GetThaJoltage() int {
	resultString := ""

	for _, value := range bank.TheChosenTwelve {
		resultString += strconv.Itoa(value)
	}
	resultInt, _ := strconv.Atoi(resultString)
	return resultInt
}
