package local

import "github.com/shytikov/advent-of-go/shared"

type Octopus shared.Node

func (o *Octopus) increaseEnergy() int {
	status := o.Value.(*Status)
	status.position.Z = (status.position.Z + 1) % 10

	return status.position.Z
}

func (o *Octopus) getEnergy() int {
	status := o.Value.(*Status)

	return status.position.Z
}

func (o *Octopus) setFlashed(value bool) {
	status := o.Value.(*Status)
	status.flashed = value
}

func (o *Octopus) getFlashed() bool {
	status := o.Value.(*Status)

	return status.flashed
}

func (o *Octopus) getPosition() (int, int, int) {
	status := o.Value.(*Status)

	return status.position.X, status.position.Y, status.position.Z
}

func (o *Octopus) charge() {
	if !o.getFlashed() {
		if o.increaseEnergy() == 0 {
			o.setFlashed(true)
			for _, current := range o.Links {
				(*Octopus)(current).charge()
			}
		}
	}
}
