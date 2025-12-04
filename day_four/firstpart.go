package day_four

import (
	"AoC25/helper"
	"fmt"
)

func SolveThaThingPart2() {
	lines, err := helper.ReadLinesAsStringArrays("day_four/input.txt")

	if err != nil {
		return
	}

	theStorage := new(Storage)
	theStorage.Field = lines
	theStorage.ResultField = helper.DeepCopy(lines)

	for y := 0; y < len(theStorage.Field[0]); y++ {
		for x := 0; x < len(theStorage.Field); x++ {
			count := 0

			if theStorage.Field[x][y] != "@" {
				continue
			}

			if theStorage.CheckTopLeft(x, y) {
				count++
			}
			if theStorage.CheckTop(x, y) {
				count++
			}
			if theStorage.CheckTopRight(x, y) {
				count++
			}
			if theStorage.CheckRight(x, y) {
				count++
			}
			if theStorage.CheckBottomRight(x, y) {
				count++
			}
			if theStorage.CheckBottom(x, y) {
				count++
			}
			if theStorage.CheckBottomLeft(x, y) {
				count++
			}
			if theStorage.CheckLeft(x, y) {
				count++
			}

			if count < 4 {
				theStorage.ResultField[x][y] = "H"
			}
		}
	}

	for line := 0; line < len(theStorage.Field[0]); line++ {
		fmt.Println(theStorage.ResultField[line])
	}

	fmt.Println(theStorage.GetTheTowlCount())
}
