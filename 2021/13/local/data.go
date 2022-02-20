package local

import (
	"io/ioutil"
	"strings"

	"github.com/shytikov/advent-of-go/shared"
)

type Data struct {
	Points []shared.Point
	Folds  []int
}

func Read(filename string) (result Data) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	return parseData(string(content))
}

func parseData(content string) (result Data) {
	chunks := strings.Split(trimLines(content), "\n\n")

	points := make(chan []shared.Point)
	// folds := make(chan []int)

	getPoints(chunks[0], points)

	result.Points = <-points

	return
}

func trimLines(content string) string {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return strings.Join(lines, "\n")
}

func getPoints(content string, result chan []shared.Point) {
	points := make([]shared.Point, len(content))

	for i, line := range strings.Split(content, "\n") {
		points[i] = shared.ParsePoint(line)
	}

	result <- points
}
