package local

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/shytikov/advent-of-go/shared"
)

type Data struct {
	Points []shared.Point
	Folds  []shared.Point
}

func Read(filename string) (result Data) {
	content, err := ioutil.ReadFile(filename)
	shared.ActOn(err)

	return parseData(string(content))
}

func parseData(content string) (result Data) {
	chunks := strings.Split(trimLines(content), "\n\n")

	points := make(chan []shared.Point)
	folds := make(chan []shared.Point)

	go getPoints(chunks[0], points)
	go getFolds(chunks[1], folds)

	result.Points = <-points
	result.Folds = <-folds

	return
}

func trimLines(content string) string {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func getPoints(content string, result chan []shared.Point) {
	lines := strings.Split(content, "\n")
	points := make([]shared.Point, len(lines))

	for i, line := range lines {
		points[i] = shared.ParsePoint(line)
	}

	result <- points
}

func getFolds(content string, result chan []shared.Point) {
	lines := strings.Split(content, "\n")
	folds := make([]shared.Point, len(lines))

	for i, line := range lines {
		chunks := strings.Split(line, "=")
		value, _ := strconv.Atoi(chunks[1])
		if strings.HasSuffix(chunks[0], "x") {
			folds[i] = shared.Point{X: value, Y: 0}
		} else {
			folds[i] = shared.Point{X: 0, Y: value}
		}
	}

	result <- folds
}
