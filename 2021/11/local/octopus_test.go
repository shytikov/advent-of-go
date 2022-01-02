package local

import (
	"testing"

	"github.com/shytikov/advent-of-go/shared"
)

func TestOctopusIncreaseEnergy(t *testing.T) {
	// Arrange
	actual := Octopus{
		position: &shared.Point{
			X: 0,
			Y: 0,
			Z: 9,
		},
	}

	expected := 0

	// Act
	for i := 0; i < 10; i++ {
		actual.increaseEnergy()
	}

	// Assert
	if actual.position.Z != expected {
		t.Errorf("Energy level was incorrect, got: %v, want: %v", actual.position.Z, expected)
	}
}

func TestOctopusCharge(t *testing.T) {
	// Arrange
	cavern := make(Cavern, 3)
	expected := []int{0, 0, 1}

	// It should lit up fight after first step because:
	// * it's about to lit up
	cavern[0] = Octopus{
		position: &shared.Point{
			X: 0,
			Y: 0,
			Z: 9,
		},
		neighbours: []*Octopus{
			&cavern[1],
			&cavern[2],
		},
	}

	// It should lit up after first step, because:
	// * node 0 will pass its charge
	// * node 2 will pass its charge
	cavern[1] = Octopus{
		position: &shared.Point{
			X: 1,
			Y: 1,
			Z: 8,
		},
	}

	// It should lit up right afer first step, because:
	// * it's about to lit up
	// * node 0 will pass its charge
	cavern[2] = Octopus{
		position: &shared.Point{
			X: 2,
			Y: 2,
			Z: 9,
		},
		neighbours: []*Octopus{
			&cavern[0], // Adding circular dependency
			&cavern[1],
		},
	}

	// Act
	cavern[2].charge()
	actual := []int{
		cavern[0].position.Z,
		cavern[1].position.Z,
		cavern[2].position.Z,
	}

	// Assert
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Result was incorrect, got: %v, want: %v", actual[i], expected[i])
			return
		}
	}
}
