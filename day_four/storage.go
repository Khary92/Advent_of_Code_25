package day_four

type Storage struct {
	Field       [][]string
	ResultField [][]string
	Target      int
}

func (storage *Storage) GetTheTowlCount() int {
	count := 0

	for y := 0; y < len(storage.Field[0]); y++ {
		for x := 0; x < len(storage.Field); x++ {
			if storage.ResultField[x][y] == "H" {
				count++
			}
		}
	}
	return count
}

func (storage *Storage) CheckTopLeft(xCoord int, yCoord int) bool {
	if yCoord-1 < 0 {
		return false
	}

	if xCoord-1 < 0 {
		return false
	}

	if storage.Field[xCoord-1][yCoord-1] == "@" {
		return true
	}

	return false
}

func (storage *Storage) CheckTop(xCoord int, yCoord int) bool {
	if yCoord-1 < 0 {
		return false
	}

	if storage.Field[xCoord][yCoord-1] == "@" {
		return true
	}

	return false
}

func (storage *Storage) CheckTopRight(xCoord int, yCoord int) bool {
	if yCoord-1 < 0 {
		return false
	}

	if xCoord+1 >= len(storage.Field) {
		return false
	}

	if storage.Field[xCoord+1][yCoord-1] == "@" {
		return true
	}

	return false
}

func (storage *Storage) CheckRight(xCoord int, yCoord int) bool {
	if xCoord+1 >= len(storage.Field) {
		return false
	}

	if storage.Field[xCoord+1][yCoord] == "@" {
		return true
	}

	return false
}

func (storage *Storage) CheckBottomRight(xCoord int, yCoord int) bool {
	if yCoord+1 >= len(storage.Field[0]) {
		return false
	}

	if xCoord+1 >= len(storage.Field) {
		return false
	}

	if storage.Field[xCoord+1][yCoord+1] == "@" {
		return true
	}

	return false
}

func (storage *Storage) CheckBottom(xCoord int, yCoord int) bool {
	if yCoord+1 >= len(storage.Field[0]) {
		return false
	}

	if storage.Field[xCoord][yCoord+1] == "@" {
		return true
	}

	return false
}

func (storage *Storage) CheckBottomLeft(xCoord int, yCoord int) bool {
	if yCoord+1 >= len(storage.Field[0]) {
		return false
	}

	if xCoord-1 < 0 {
		return false
	}

	if storage.Field[xCoord-1][yCoord+1] == "@" {
		return true
	}

	return false
}

func (storage *Storage) CheckLeft(xCoord int, yCoord int) bool {
	if xCoord-1 < 0 {
		return false
	}

	if storage.Field[xCoord-1][yCoord] == "@" {
		return true
	}

	return false
}
