package day_five

type Range struct {
	Lower int
	Upper int
	Size  int
}

type RangeHolder struct {
	CurrentRanges []Range
	NewRanges     []Range
}

func (rangeValue *Range) CalcRange() {
	rangeValue.Size = rangeValue.Upper - rangeValue.Lower
}

func (rnage *Range) ContainsRange(anotherRnage Range) bool {
	if rnage.Lower <= anotherRnage.Lower && rnage.Upper >= anotherRnage.Upper {
		return true
	}

	return false
}
