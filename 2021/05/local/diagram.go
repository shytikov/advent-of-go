package local

type Diagram [][]int

func (d Diagram) Draw(v Vent) {
	var beginX, beginY, endX, endY int

	if v.IsHorizontal() {
		if v.From.X > v.To.X {
			beginX, endX = v.To.X, v.From.X
		} else {
			beginX, endX = v.From.X, v.To.X
		}

		beginY, endY = v.From.Y, v.From.Y
	}

	if v.IsVertical() {
		if v.From.Y > v.To.Y {
			beginY, endY = v.To.Y, v.From.Y
		} else {
			beginY, endY = v.From.Y, v.To.Y
		}

		beginX, endX = v.From.X, v.From.X
	}

	for i := beginY; i <= endY; i++ {
		for j := beginX; j <= endX; j++ {
			d[i][j] += 1
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
