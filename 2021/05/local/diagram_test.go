package local

import (
	"testing"
)

func TestDiagramDraw(t *testing.T) {
	// Arrange
	definition :=
		`0,9 -> 5,9
		8,0 -> 0,8
		9,4 -> 3,4
		2,2 -> 2,1
		7,0 -> 7,4
		6,4 -> 2,0
		0,9 -> 2,9
		3,4 -> 1,4
		0,0 -> 8,8
		5,5 -> 8,2`

	data := parseData(definition)
	actual := data.CreateDiagram()

	expected := Diagram{
		{1, 0, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0, 0, 0, 2, 0, 0},
		{0, 0, 2, 0, 1, 0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 2, 0, 2, 0, 0},
		{0, 1, 1, 2, 3, 1, 3, 2, 1, 1},
		{0, 0, 0, 1, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	}

	// Act
	for _, vent := range data.Vents {
		if vent.IsOrthogonal() {
			actual.Draw(vent)

			//			fmt.Println()
			//
			//			for _, row := range actual {
			//				fmt.Println(row)
			//			}
		}

	}

	// Assert
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if actual[i][j] != expected[i][j] {
				t.Errorf("Value on (%d, %d) was incorrect, got: %d, want: %d", i, j, actual[i][j], expected[i][j])
				return
			}
		}
	}
}

func TestDiagramCount(t *testing.T) {
	// Arrange
	diagram := Diagram{
		{1, 0, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0, 0, 0, 2, 0, 0},
		{0, 0, 2, 0, 1, 0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 2, 0, 2, 0, 0},
		{0, 1, 1, 2, 3, 1, 3, 2, 1, 1},
		{0, 0, 0, 1, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	}

	expected := 12

	// Act
	actual := diagram.CountGreaterThan(1)

	// Assert
	if actual != expected {
		t.Errorf("Count was incorrect, got: %d, want: %d", actual, expected)
	}
}

func TestDiagramDrawOrthogonal(t *testing.T) {
	// Arrange
	definition :=
		`0,9 -> 5,9
		7,0 -> 7,4`

	data := parseData(definition)
	actual := data.CreateDiagram()
	expected := Diagram{
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 0, 0},
	}

	// Act
	for _, vent := range data.Vents {
		if vent.IsOrthogonal() {
			actual.Draw(vent)
		}
	}

	// Assert
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if actual[i][j] != expected[i][j] {
				t.Errorf("Value on (%d, %d) was incorrect, got: %d, want: %d", i, j, actual[i][j], expected[i][j])
				return
			}
		}
	}
}
