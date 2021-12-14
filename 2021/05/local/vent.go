package local

type Vent struct {
	From Point
	To   Point
}

type Point struct {
	X int
	Y int
}

func (v Vent) IsOrthogonal() bool {
	return v.From.X == v.To.X || v.From.Y == v.To.Y
}

func (v Vent) IsDiagonal() bool {
	return !v.IsOrthogonal()
}
