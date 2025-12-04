package day_four

import (
	"AoC25/helper"
	"fmt"
)

func SolveThaThing() {
	lines, err := helper.ReadLinesAsStringArrays("day_four/input.txt")

	if err != nil {
		return
	}

	theStorage := new(Storage)
	theStorage.Field = lines
	theStorage.ResultField = helper.DeepCopy(lines)

	lastRemovedThingies := 0

	for {
		currentRemovedThingies := theStorage.DoTheTakingThing()

		if currentRemovedThingies == lastRemovedThingies {
			break
		}

		lastRemovedThingies = currentRemovedThingies
	}

	fmt.Println("Removed thingies in total: ", lastRemovedThingies)
}

func (storage *Storage) DoTheTakingThing() (takenCount int) {
	storage.Field = storage.ResultField
	storage.ResultField = helper.DeepCopy(storage.ResultField)

	for y := 0; y < len(storage.Field[0]); y++ {
		for x := 0; x < len(storage.Field); x++ {
			count := 0

			if storage.Field[x][y] != "@" {
				continue
			}

			if storage.CheckTopLeft(x, y) {
				count++
			}
			if storage.CheckTop(x, y) {
				count++
			}
			if storage.CheckTopRight(x, y) {
				count++
			}
			if storage.CheckRight(x, y) {
				count++
			}
			if storage.CheckBottomRight(x, y) {
				count++
			}
			if storage.CheckBottom(x, y) {
				count++
			}
			if storage.CheckBottomLeft(x, y) {
				count++
			}
			if storage.CheckLeft(x, y) {
				count++
			}

			if count < 4 {
				storage.ResultField[x][y] = "H"
			}
		}
	}

	fmt.Println("------------------------------------------------------")
	for line := 0; line < len(storage.Field[0]); line++ {
		fmt.Println(storage.ResultField[line])
	}

	takenCount = storage.GetTheTowlCount()
	return
}
