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

	x := make(chan int, count)
	y := make(chan int, count)

	result.Vents = make([]Vent, count)
	result.Min = Point{0, 0}
	result.Max = Point{0, 0}

	var wg sync.WaitGroup

	for i, line := range lines {
		wg.Add(1)

		go func(index int, definition string) {
			defer wg.Done()

			vent := parseVent(definition)

			result.Vents[index] = vent

			getMaxToChannel(x, vent.From.X, vent.To.X)
			getMaxToChannel(y, vent.From.Y, vent.To.Y)
		}(i, line)
	}

	wg.Wait()
	close(x)
	close(y)

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

func getMaxToChannel(output chan int, a, b int) {
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
