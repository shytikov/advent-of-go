package local

type Vent struct {
	From Point
	To   Point
}

type Point struct {
	X int
	Y int
}

func (v Vent) IsHorizontal() bool {
	return v.From.X == v.To.X
}

func (v Vent) IsVertical() bool {
	return v.From.Y == v.To.Y
}

func (v Vent) IsOrthogonal() bool {
	return v.IsHorizontal() || v.IsVertical()
}

func (v Vent) IsDiagonal() bool {
	return !v.IsOrthogonal()
}
