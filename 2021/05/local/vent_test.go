package local

import (
	"testing"

	"github.com/shytikov/advent-of-go/shared"
)

func TestIsOrthogonal(t *testing.T) {
	// Arrange
	vent := Vent{
		From: shared.Point{
			X: 599,
			Y: 531,
		},
		To: shared.Point{
			X: 599,
			Y: 32,
		},
	}

	expected := true

	// Act
	actual := vent.IsOrthogonal()

	// Assert
	if actual != expected {
		t.Errorf("Result was incorrect, got: %t, want: %t", actual, expected)
	}
}
