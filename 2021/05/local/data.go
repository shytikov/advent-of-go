package local

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Data struct {
	Vents []Vent
}

type Vent struct {
	From Point
	To   Point
}

type Point struct {
	X int
	Y int
}

func Read(filename string) (result Data) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	result.Vents = make([]Vent, len(lines))

	for i, line := range lines {
		go func(index int, definition string) {
			result.Vents[index] = parseVent(definition)
		}(i, line)
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
		From: parsePoint(chunks[0]),
		To:   parsePoint(chunks[1]),
	}
}
