package shared

import (
	"strconv"
	"strings"
)

// Point handy data structure to represent a point in two-dimensional space
type Point struct {
	X int
	Y int
	Z int
}

func ParsePoint(definition string) (result Point) {
	chunks := strings.Split(definition, ",")

	result.X, _ = strconv.Atoi(chunks[0])
	result.Y, _ = strconv.Atoi(chunks[1])

	return
}
