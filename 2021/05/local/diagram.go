package local

type Diagram [][]int

func (d Diagram) Draw(v Vent) {
	if v.IsHorizontal() {
		var begin int
		var end int

		if v.From.X > v.To.X {
			begin = v.To.X
			end = v.From.X
		} else {
			begin = v.From.X
			end = v.To.X
		}

		y := v.From.Y

		for i := begin; i <= end; i++ {
			d[y][i] = d[y][i] + 1
		}
	}

	if v.IsVertical() {
		var begin int
		var end int

		if v.From.Y > v.To.Y {
			begin = v.To.Y
			end = v.From.Y
		} else {
			begin = v.From.Y
			end = v.To.Y
		}

		x := v.From.X

		for i := begin; i <= end; i++ {
			d[i][x] = d[i][x] + 1
		}
	}
}
