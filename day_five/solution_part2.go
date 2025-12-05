package day_five

import "fmt"

func SolveThaThingPart2() {
	_, ranges := SeparateTheThing()

	thaHodler := RangeHolder{}
	thaHodler.CurrentRanges = ranges
	thaHodler.NewRanges = ranges

	thaHodler.SomethingRecursiveIGuess()
}

func (hodler *RangeHolder) SomethingRecursiveIGuess() {

	for {
		hodler.CurrentRanges = DeepCopy(hodler.NewRanges)
		hodler.NewRanges = make([]Range, 0)

		for i := 0; i < len(hodler.CurrentRanges); i++ {

			isContainedInAnotherRange := false

			for j := 0; j < len(hodler.CurrentRanges); j++ {

				if i == j {
					continue
				}

				if hodler.CurrentRanges[j].ContainsRange(hodler.CurrentRanges[i]) {
					isContainedInAnotherRange = true
					break
				}
			}

			if !isContainedInAnotherRange {
				hodler.NewRanges = append(hodler.NewRanges, hodler.CurrentRanges[i])
			}

			isContainedInAnotherRange = false
		}

		if len(hodler.CurrentRanges) == len(hodler.NewRanges) {
			break
		}

		fmt.Printf("Tha fresh count is %d\n", len(hodler.NewRanges))
	}
}

func DeepCopy(field []Range) []Range {
	copyField := make([]Range, len(field))
	copy(copyField, field)
	return copyField
}
