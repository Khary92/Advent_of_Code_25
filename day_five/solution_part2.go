package day_five

import (
	"fmt"
	"slices"
)

func SolveThaThingPart2() {
	_, ranges := SeparateTheThing()

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.Lower < b.Lower {
			return -1
		} else if a.Lower > b.Lower {
			return 1
		}
		return 0
	})

	mergedRanges := MergeAthingOrSO(ranges)

	fmt.Println(mergedRanges)

	count := 0
	for i := 0; i < len(mergedRanges); i++ {
		mergedRanges[i].CalcRange()
		count += mergedRanges[i].Size
	}

	fmt.Println(count)
}

func MergeAthingOrSO(ranges []Range) []Range {
	result := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		last := &result[len(result)-1]
		curr := ranges[i]

		if curr.Lower <= last.Upper {
			if curr.Upper > last.Upper {
				last.Upper = curr.Upper
			}
			continue
		}

		result = append(result, curr)
	}

	return result
}
