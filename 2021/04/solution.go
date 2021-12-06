package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := readBingo("./input.txt")

	if input.Draw != nil {
		resultA := make(chan int)
		resultB := make(chan int)

		go solvePuzzleA(input, resultA)
		go solvePuzzleB(input, resultB)

		fmt.Println(<-resultA)
		fmt.Println(<-resultB)
	} else {
		fmt.Errorf("failure when reading input data")
	}
}

func solvePuzzleA(input Data, result chan int) {
	result <- 0
}

func solvePuzzleB(input Data, result chan int) {
	result <- 0
}

type Board struct {
	Initial  [][]int
	Marked   [][]int
	Unmarked [][]int
}

type Data struct {
	Draw   []int
	Boards []Board
}

func readBingo(filename string) Data {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return Data{}
	}

	lines := strings.Split(string(content), "\n")

	result := Data{
		Draw: getDraw(lines),
	}

	return result
}

func getDraw(lines []string) (result []int) {
	draw := strings.Split(lines[0], ",")
	result = make([]int, len(draw))

	for i, number := range draw {
		value, err := strconv.Atoi(number)

		if err != nil {
			continue
		}
		result[i] = value
	}

	return result
}
