package local

import (
	"github.com/shytikov/advent-of-go/shared"
)

type Octopus struct {
	position   shared.Point
	neighbours Area
	flashed    bool
}

func (o *Octopus) increaseEnergy() {
	o.position.Z = (o.position.Z + 1) % 10
}

func (o *Octopus) charge() {
	if !o.flashed {
		o.increaseEnergy()
		if o.position.Z == 0 && !o.flashed {
			o.flashed = true
			for _, link := range o.neighbours {
				link.charge()
			}
		}
	}
}
