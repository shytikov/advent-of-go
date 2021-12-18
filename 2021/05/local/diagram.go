package local

type Diagram [][]int

func (d Diagram) Draw(vent Vent) {
	if vent.IsOrthogonal() {
		d.drawOrthogonal(vent)
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

func (d Diagram) drawOrthogonal(vent Vent) {
	if vent.From.X > vent.To.X || vent.From.Y > vent.To.Y {
		vent.From, vent.To = vent.To, vent.From
	}

	for i := vent.From.Y; i <= vent.To.Y; i++ {
		for j := vent.From.X; j <= vent.To.X; j++ {
			d[i][j] += 1
		}
	}
}
