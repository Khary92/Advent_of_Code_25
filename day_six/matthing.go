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

func DoTheRightStructure(input []string) (hodlers []MathHolder) {
	// utter crap in this language. This is partly AI generated code as datastructures still are pure garbage to handle without intricate helpers...
	// The idea is simple.
	// 1. Make the Lines into a matrix by character
	// 2. Rotate the matrix 90 degrees to the left
	// 3. Scan the horizontal lines for the numbers
	// 4. If it is not a number, nor an empty space, it is the operation
	// 5. If the line consists of only spaces, the next iteration starts

	theSlicedThang := make([][]string, 0, len(input))
	for _, value := range input {
		theSlicedThang = append(theSlicedThang, helper.StringToSlice(value))
	}

	thaMatrix := RotateTheCrapToTheLEft(theSlicedThang)
	hodlers = MakeThaHodlersWithThisCrap(thaMatrix)

	return
}

func MakeThaHodlersWithThisCrap(matrix [][]string) (hodlers []MathHolder) {
	var currentHodler MathHolder

	for _, row := range matrix {
		isEmpty := true
		for _, ch := range row {
			if ch != " " {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			if len(currentHodler.Numbers) > 0 || currentHodler.Operation != "" {
				hodlers = append(hodlers, currentHodler)
				currentHodler = MathHolder{}
			}
			continue
		}

		var numStr string
		for _, ch := range row {
			if ch >= "0" && ch <= "9" {
				numStr += ch
			} else if ch == "+" || ch == "*" {
				currentHodler.Operation = ch
			}
		}

		if numStr != "" {
			number, err := strconv.Atoi(numStr)
			if err == nil {
				currentHodler.Numbers = append(currentHodler.Numbers, number)
			}
		}
	}

	if len(currentHodler.Numbers) > 0 || currentHodler.Operation != "" {
		hodlers = append(hodlers, currentHodler)
	}

	return hodlers
}

func RotateTheCrapToTheLEft(matrix [][]string) [][]string {
	if len(matrix) == 0 {
		return nil
	}
	rows := len(matrix)
	cols := len(matrix[0])

	rotated := make([][]string, cols)
	for i := 0; i < cols; i++ {
		rotated[i] = make([]string, rows)
		for j := 0; j < rows; j++ {
			rotated[i][j] = matrix[j][cols-1-i]
		}
	}

	for _, row := range rotated {
		fmt.Println(row)
	}

	return rotated
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
