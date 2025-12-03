package day_three

import "strconv"

type ValueHolder struct {
	FirstBiggestNumber  int
	SecondBiggestNumber int
	FirstIndex          int
	SecondIndex         int
	ValueArray          []int
}

func CalculateThaThing(values []int) (result ValueHolder) {
	result.ValueArray = values
	result.FirstBiggestNumber = values[0]

	for i := 1; i < len(values)-1; i++ {
		if values[i] > result.FirstBiggestNumber {
			result.FirstBiggestNumber = values[i]
			result.FirstIndex = i
		}
	}

	result.SetSecondBiggestNumber()

	return result
}

func (hodler *ValueHolder) SetSecondBiggestNumber() {
	hodler.SecondBiggestNumber = hodler.ValueArray[hodler.FirstIndex+1]
	hodler.SecondIndex = hodler.ValueArray[hodler.FirstIndex+1]
	for i := hodler.FirstIndex + 1; i < len(hodler.ValueArray); i++ {
		if hodler.ValueArray[i] > hodler.SecondBiggestNumber {
			hodler.SecondBiggestNumber = hodler.ValueArray[i]
			hodler.SecondIndex = i
		}
	}
}

func (hodler *ValueHolder) GetThaJoltage() int {
	resultString := strconv.Itoa(hodler.FirstBiggestNumber) + strconv.Itoa(hodler.SecondBiggestNumber)
	resultInt, _ := strconv.Atoi(resultString)
	return resultInt
}
