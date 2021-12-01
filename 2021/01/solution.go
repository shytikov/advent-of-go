package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(first())
}

func first() int {
	// Counter for increased depth measurements
	counter := 0

	// Previous depth value – intentionally set to big number
	previous := math.MaxInt16
	for _, measurement := range readInput("input.txt") {
		current, err := strconv.Atoi(measurement)

		if err != nil {
			panic(fmt.Errorf("cannot convert measurement to number"))
		}

		if current-previous > 0 {
			counter++
		}

		previous = current
	}

	return counter
}

func readInput(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("cannot find input file"))
	}

	return strings.Split(string(content), "\n")
}
