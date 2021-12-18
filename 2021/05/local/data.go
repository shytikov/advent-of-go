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

		result.Max.X = getMaxFromChannel(x)
	}()

	go func() {
		defer wg.Done()

		result.Max.Y = getMaxFromChannel(y)
	}()

	wg.Wait()

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

func setMaxToChannel(output chan int, a, b int) {
	if a >= b {
		output <- a
	} else {
		output <- b
	}
}

func getMaxFromChannel(input chan int) (result int) {
	for number := range input {
		if result < number {
			result = number
		}
	}

	return
}
