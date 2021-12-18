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
	diagram := data.CreateDiagram()

	// Act
	diagram.Draw(data.Vents[0])

	// Assert
	// fmt.Println(diagram)
	//	if actualHorizontal != expectedHorizontal {
	//		t.Errorf("Vent (%s) was vertical, IsHorizontal() got: %t, want: %t", vent.Definition, actualHorizontal, expectedHorizontal)
	//	}
	//
	//	if actualVertical != expectedVertical {
	//		t.Errorf("Vent (%s) was vertical, IsVertical() got: %t, want: %t", vent.Definition, actualVertical, expectedVertical)
	//	}
}
