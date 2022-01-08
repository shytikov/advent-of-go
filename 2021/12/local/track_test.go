package local

import (
	"testing"

	"github.com/shytikov/advent-of-go/shared"
)

func TestTrackAdd(t *testing.T) {
	// Arrange
	track := Track{path: "start,A,B,c,"}
	nodes := []*shared.Node{
		&shared.Node{Value: "d"},
		&shared.Node{Value: "end"},
	}

	expected := []string{
		"start,A,B,c,d",
		"start,A,B,c,d,end",
	}

	// Act
	actual := []string{
		track.Add(nodes[0]),
		track.Add(nodes[1]),
	}

	// Assert
	for i, value := range expected {
		if value != actual[i] {
			t.Errorf("Result was incorrect, got: %v, want: %v",
				value,
				actual[i])
			return
		}
	}
}

func TestTrackVisited(t *testing.T) {
	// Arrange
	expected := map[string]int{
		"start": 1,
		"end":   1,
		"A":     3,
		"B":     1,
		"c":     2,
		"d":     1,
	}

	// Act
	actual := Track{}
	actual.Add(&shared.Node{Value: "start"})
	actual.Add(&shared.Node{Value: "A"})
	actual.Add(&shared.Node{Value: "A"})
	actual.Add(&shared.Node{Value: "A"})
	actual.Add(&shared.Node{Value: "B"})
	actual.Add(&shared.Node{Value: "c"})
	actual.Add(&shared.Node{Value: "c"})
	actual.Add(&shared.Node{Value: "d"})
	actual.Add(&shared.Node{Value: "end"})

	// Assert
	for i, value := range expected {
		if value != actual.Visited[i] {
			t.Errorf("Result was incorrect, got: %v, want: %v",
				actual.Visited,
				expected)
			return
		}
	}
}
