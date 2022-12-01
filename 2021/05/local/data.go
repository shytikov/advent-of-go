package local

import (
	"os"
	"strings"
	"sync"

	"github.com/shytikov/advent-of-go/shared"
)

type Data struct {
	Vents []Vent
	Min   shared.Point
	Max   shared.Point
}

func Read(filename string) Data {
	content, err := os.ReadFile(filename)
	shared.ActOn(err)

	return parseData(string(content))
}

func (d Data) CreateDiagram() (result Diagram) {
	// Preventing off by one error
	lenX := d.Max.X + 1
	lenY := d.Max.Y + 1

	// Vertical coordinate (Y) should come first
	result = make(Diagram, lenY)

	for i := range result {
		result[i] = make([]int, lenX)
	}

	return
}

func parseData(content string) (result Data) {
	lines := strings.Split(content, "\n")
	count := len(lines)

	x := make(chan int, count)
	y := make(chan int, count)

	result.Vents = make([]Vent, count)

	var wg sync.WaitGroup

	// Parsing all vents at once
	for i, line := range lines {
		wg.Add(1)

		go func(index int, definition string) {
			defer wg.Done()

			vent := parseVent(definition)

			setMaxToChannel(x, vent.From.X, vent.To.X)
			setMaxToChannel(y, vent.From.Y, vent.To.Y)

			result.Vents[index] = vent
		}(i, line)
	}

	wg.Wait()
	close(x)
	close(y)

	// Calculating max coordinates in parallel as well
	wg.Add(2)

	go func() {
		defer wg.Done()

		result.Max.X = shared.GetMaxFromChannel(x)
	}()

	go func() {
		defer wg.Done()

		result.Max.Y = shared.GetMaxFromChannel(y)
	}()

	wg.Wait()

	return
}

func parseVent(definition string) (result Vent) {
	definition = strings.TrimSpace(definition)
	chunks := strings.Split(definition, " -> ")

	from := shared.ParsePoint(chunks[0])
	to := shared.ParsePoint(chunks[1])

	result = Vent{
		From:   from,
		To:     to,
		Vector: getDirection(from, to),
	}

	return
}

func getDirection(from, to shared.Point) (result shared.Point) {
	result = shared.Point{X: 1, Y: 1}

	if from.X > to.X {
		result.X = -1
	}
	if from.Y > to.Y {
		result.Y = -1
	}
	if from.X == to.X {
		result.X = 0
	}
	if from.Y == to.Y {
		result.Y = 0
	}

	return
}

func setMaxToChannel(output chan int, a, b int) {
	if a >= b {
		output <- a
	} else {
		output <- b
	}
}
