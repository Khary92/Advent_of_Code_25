package day_six

import (
	"AoC25/helper"
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
