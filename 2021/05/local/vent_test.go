package local

import "testing"

func TestIsOrthogonal(t *testing.T) {
	// Arrange
	vent := Vent{
		Definition: "599,531 -> 599,32",
		From: Point{
			599,
			531,
		},
		To: Point{
			599,
			32,
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

func TestVentIsVertical(t *testing.T) {
	// Arrange
	vent := Vent{
		Definition: "771,406 -> 120,406",
		From: Point{
			771,
			406,
		},
		To: Point{
			120,
			406,
		},
	}

	expectedHorizontal := false
	expectedVertical := true

	// Act
	actualHorizontal := vent.IsHorizontal()
	actualVertical := vent.IsVertical()

	// Assert
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Vent (%s) was vertical, IsHorizontal() got: %t, want: %t", vent.Definition, actualHorizontal, expectedHorizontal)
	}

	if actualVertical != expectedVertical {
		t.Errorf("Vent (%s) was vertical, IsVertical() got: %t, want: %t", vent.Definition, actualVertical, expectedVertical)
	}
}

func TestVentIsHorizontal(t *testing.T) {
	// Arrange
	vent := Vent{
		Definition: "599,531 -> 599,32",
		From: Point{
			599,
			531,
		},
		To: Point{
			599,
			32,
		},
	}

	expectedHorizontal := true
	expectedVertical := false

	// Act
	actualHorizontal := vent.IsHorizontal()
	actualVertical := vent.IsVertical()

	// Assert
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Vent (%s) was horizontal, IsHorizontal() got: %t, want: %t", vent.Definition, actualHorizontal, expectedHorizontal)
	}

	if actualVertical != expectedVertical {
		t.Errorf("Vent (%s) was horizontal, IsVertical() got: %t, want: %t", vent.Definition, actualVertical, expectedVertical)
	}
}
