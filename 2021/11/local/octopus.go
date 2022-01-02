package local

import (
	"github.com/shytikov/advent-of-go/shared"
)

type Octopus struct {
	position   *shared.Point
	neighbours []*Octopus
	flashed    bool
}

func (o *Octopus) increaseEnergy() {
	if o.position.Z == 9 {
		o.position.Z = 0
	} else {
		o.position.Z += 1
	}
}

func (o *Octopus) charge() {
	if !o.flashed {
		o.position.Z = (o.position.Z + 1) % 10
		if o.position.Z == 0 {
			o.flashed = true
			for _, link := range o.neighbours {
				link.charge()
			}
		}
	}
}
