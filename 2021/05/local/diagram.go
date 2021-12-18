package local

type Diagram [][]int

func (d Diagram) Draw(vent Vent) {
	current := vent.From

	for {
		d[current.Y][current.X] += 1

		if current.X == vent.To.X && current.Y == vent.To.Y {
			break
		} else {
			current.X += vent.Vector.X
			current.Y += vent.Vector.Y
		}
	}
}

func (d Diagram) CountGreaterThan(threshold int) (result int) {
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[i]); j++ {
			if d[i][j] > threshold {
				result++
			}
		}
	}

	return
}
