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
	length := len(chunks)

	if length >= 1 {
		result.X, _ = strconv.Atoi(chunks[0])
	}
	if length >= 2 {
		result.Y, _ = strconv.Atoi(chunks[1])
	}
	if length >= 3 {
		result.Z, _ = strconv.Atoi(chunks[2])
	}

	return
}
