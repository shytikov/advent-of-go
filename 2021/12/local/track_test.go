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

func TestTrackGetStats(t *testing.T) {
	// Arrange
	track := Track{}
	track.Add(&shared.Node{Value: "start"})
	track.Add(&shared.Node{Value: "A"})
	track.Add(&shared.Node{Value: "A"})
	track.Add(&shared.Node{Value: "A"})
	track.Add(&shared.Node{Value: "B"})
	track.Add(&shared.Node{Value: "c"})
	track.Add(&shared.Node{Value: "c"})
	track.Add(&shared.Node{Value: "d"})
	track.Add(&shared.Node{Value: "end"})

	expected := []int{1, 1, 3, 2}

	// Act
	actual := make([]int, 4)
	actual[0], actual[1], actual[2], actual[3] = track.GetStats()

	// Assert
	for i, value := range expected {
		if value != actual[i] {
			t.Errorf("Result was incorrect, got: %v, want: %v",
				actual,
				expected)
			return
		}
	}
}
