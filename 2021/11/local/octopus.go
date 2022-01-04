package local

import (
	"github.com/shytikov/advent-of-go/shared"
)

type Octopus struct {
	Value   shared.Point
	Links   Area
	flashed bool
}

func (o *Octopus) increaseEnergy() {
	o.Value.Z = (o.Value.Z + 1) % 10
}

func (o *Octopus) charge() {
	if !o.flashed {
		o.increaseEnergy()
		if o.Value.Z == 0 && !o.flashed {
			o.flashed = true
			for _, link := range o.Links {
				link.charge()
			}
		}
	}
}
