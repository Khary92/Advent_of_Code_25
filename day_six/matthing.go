package day_six

import (
	"AoC25/helper"
	"fmt"
	"strconv"
)

type MathHolder struct {
	Numbers   []int
	Operation string
	Result    int
}

func DoStructure(input []string) []MathHolder {
	theThing := make([][]string, 0, len(input))
	for _, value := range input {
		theThing = append(theThing, helper.SplitByChars(value, " "))
	}

	hodlers := make([]MathHolder, 0, len(theThing))
	for i := 0; i < len(theThing[0]); i++ {
		hodlerValues := make([]string, 0)
		for j := 0; j < len(theThing); j++ {
			hodlerValues = append(hodlerValues, theThing[j][i])
		}

		singleHodler := MathHolder{}
		for j := 0; j < len(hodlerValues); j++ {
			if j == len(hodlerValues)-1 {
				singleHodler.Operation = hodlerValues[j]
				continue
			}

			theNumber, err := strconv.Atoi(hodlerValues[j])
			if err == nil {
				singleHodler.Numbers = append(singleHodler.Numbers, theNumber)
			}

		}

		hodlers = append(hodlers, singleHodler)
	}

	return hodlers
}

func DoTheRightStructure(input []string) []MathHolder {
	theThing := make([][]string, 0, len(input))
	for _, value := range input {
		theThing = append(theThing, helper.SplitByChars(value, " "))
	}

	hodlers := make([]MathHolder, 0, len(theThing))
	for i := 0; i < len(theThing[0]); i++ {
		hodlerValues := make([]string, 0)
		for j := 0; j < len(theThing); j++ {
			hodlerValues = append(hodlerValues, theThing[j][i])
		}

		GetTheMathHodlerByOtherReadingThing(hodlerValues)
	}

	return hodlers
}

func GetTheMathHodlerByOtherReadingThing(input []string) (thaHodler MathHolder) {
	//get tha longest hodlerValue
	thaLongestCharsInAValueThing := 0
	// Dis always the same
	theYValue := 5

	for _, value := range input {
		if len(value) > thaLongestCharsInAValueThing {
			thaLongestCharsInAValueThing = len(value)
		}
	}

	thaArray := make([][]string, thaLongestCharsInAValueThing, theYValue)
	//Initialize the values map liek dis for thaLongestCharsInAValueThing = 4
	// [2, 3, 4, 5]
	// [3, 4, 5, -]
	// [4, 5, -, -]
	// [5, -, -, -]
	// [*, -, -, -]

	for i := 0; i < len(input); i++ {
		chunkifiedNumberOrOperator := helper.StringToSlice(input[i])

		for i := 0; i < thaLongestCharsInAValueThing; i++ {
			if i >= len(chunkifiedNumberOrOperator) {
				thaArray[i] = append(thaArray[i], "-")
				continue
			}

			thaArray[i] = append(thaArray[i], chunkifiedNumberOrOperator[i])
		}

	}

	fmt.Println(thaArray)

	//Great success in some way i guess
	//Now read the stuff and make cool hodler

	for i := 0; i < len(thaArray); i++ {
		aTheStringValue := ""
		for j := 0; j < len(thaArray[i]); j++ {
			if thaArray[i][j] == "-" {
				continue
			}

			if j == len(thaArray[i])-1 && thaArray[i][j] != "-" {
				thaHodler.Operation = thaArray[i][j]
				continue
			}

			aTheStringValue += thaArray[i][j]
		}
		thaNumber, _ := strconv.Atoi(aTheStringValue)
		thaHodler.Numbers = append(thaHodler.Numbers, thaNumber)
	}
	return
}

func (hodler *MathHolder) DoTheMath() {
	if hodler.Operation == "*" {
		hodler.Result = hodler.Numbers[0]
		for i := 1; i < len(hodler.Numbers); i++ {
			hodler.Result = hodler.Result * hodler.Numbers[i]
		}
		return
	}

	hodler.Result = hodler.Numbers[0]
	for i := 1; i < len(hodler.Numbers); i++ {
		hodler.Result = hodler.Result + hodler.Numbers[i]
	}
}
