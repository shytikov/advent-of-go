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
			d[i][y] = d[i][y] + 1
		}
	}
}
