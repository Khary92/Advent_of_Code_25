package day_one

type Dial struct {
	Value int
}

func (d *Dial) TurnLeft(count int) int {
	timesPassed := 0

	if d.Value == 0 {
		count--
		d.Value = 99
	}

	for i := 0; i < count; i++ {
		d.Value = d.Value - 1

		if d.Value == -1 {
			d.Value = 99
			timesPassed++
		}
	}

	if d.Value == 0 {
		timesPassed++
	}

	return timesPassed
}

func (d *Dial) TurnRight(count int) int {
	timesPassed := 0

	for i := 0; i < count; i++ {
		d.Value = d.Value + 1

		if d.Value == 100 {
			d.Value = 0

			if i == count-1 {
				continue
			}

			timesPassed++
		}
	}

	if d.Value == 0 {
		timesPassed++
	}

	return timesPassed
}
