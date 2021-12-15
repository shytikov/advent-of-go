package local

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Data struct {
	Vents []Vent
	Min   Point
	Max   Point
}

func Read(filename string) (result Data) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	result.Vents = make([]Vent, len(lines))

	for i, line := range lines {
		result.Vents[i] = parseVent(line)

		if result.Vents[i].From.X < result.Min.X {
			result.Min.X = result.Vents[i].From.X
		}

		if result.Vents[i].From.X > result.Max.X {
			result.Max.X = result.Vents[i].From.X
		}

		if result.Vents[i].From.Y < result.Min.Y {
			result.Min.Y = result.Vents[i].From.Y
		}

		if result.Vents[i].From.Y > result.Max.Y {
			result.Max.Y = result.Vents[i].From.Y
		}
	}

	return
}

func (d Data) CreateDiagram() (result [][]int) {
	// Preventing off by one error
	lenX := d.Max.X - d.Min.X + 1
	lenY := d.Max.Y - d.Min.Y + 1

	result = make([][]int, lenX)

	for i := range result {
		result[i] = make([]int, lenY)
	}

	return
}

func parsePoint(definition string) (result Point) {
	chunks := strings.Split(definition, ",")

	result.X, _ = strconv.Atoi(chunks[0])
	result.Y, _ = strconv.Atoi(chunks[1])

	return
}

func parseVent(definition string) Vent {
	chunks := strings.Split(definition, " -> ")

	return Vent{
		Definition: definition,
		From: parsePoint(chunks[0]),
		To:   parsePoint(chunks[1]),
	}
}
