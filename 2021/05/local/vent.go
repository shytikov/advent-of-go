package local

type Vent struct {
	From   Point
	To     Point
	Vector Point
}

type Point struct {
	X int
	Y int
}

func (v Vent) IsOrthogonal() bool {
	return v.Vector.X == 0 || v.Vector.Y == 0
}
