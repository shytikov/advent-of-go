package local

import "github.com/shytikov/advent-of-go/shared"

type Vent struct {
	From   shared.Point
	To     shared.Point
	Vector shared.Point
}

func (v Vent) IsOrthogonal() bool {
	return v.Vector.X == 0 || v.Vector.Y == 0
}
