package local

import "github.com/shytikov/advent-of-go/shared"

type Vent struct {
	From   shared.Point2D
	To     shared.Point2D
	Vector shared.Point2D
}

func (v Vent) IsOrthogonal() bool {
	return v.Vector.X == 0 || v.Vector.Y == 0
}
