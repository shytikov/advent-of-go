package local

import (
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

type Data struct {
	Vents []Vent
	Min   Point
	Max   Point
}

func Read(filename string) Data {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return Data{}
	}

	return parseData(string(content))
}

func (d Data) CreateDiagram() (result Diagram) {
	// Preventing off by one error
	lenX := d.Max.X - d.Min.X + 1
	lenY := d.Max.Y - d.Min.Y + 1

	result = make(Diagram, lenX)

	for i := range result {
		result[i] = make([]int, lenY)
	}

	return
}

func parseData(content string) (result Data) {
	lines := strings.Split(content, "\n")
	count := len(lines)

	points := make(chan Point, count*2)

	result.Vents = make([]Vent, count)

	var wg sync.WaitGroup

	for i, line := range lines {
		wg.Add(1)

		go func(index int, definition string) {
			defer wg.Done()

			result.Vents[index] = parseVent(definition)

			points <- result.Vents[index].From
			points <- result.Vents[index].To

		}(i, line)
	}

	wg.Wait()
	close(points)

	for point := range points {
		result.Min = Point{
			getMin(result.Min.X, point.X),
			getMin(result.Min.Y, point.Y),
		}

		result.Max = Point{
			getMax(result.Max.X, point.X),
			getMax(result.Max.Y, point.Y),
		}
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
	definition = strings.TrimSpace(definition)
	chunks := strings.Split(definition, " -> ")

	return Vent{
		Definition: definition,
		From:       parsePoint(chunks[0]),
		To:         parsePoint(chunks[1]),
	}
}

func getMin(previous, current int) int {
	if previous <= current {
		return previous
	} else {
		return current
	}
}

func getMax(previous, current int) int {
	if previous >= current {
		return previous
	} else {
		return current
	}
}
