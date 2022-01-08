package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/12/local"
	"github.com/shytikov/advent-of-go/shared"
)

func main() {
	input := local.Read("./input.txt")

	if input != nil {
		resultA := make(chan int)
		resultB := make(chan int)

		go solvePuzzleA(input, resultA)
		go solvePuzzleB(input, resultB)

		fmt.Println(<-resultA)
		fmt.Println(<-resultB)
	} else {
		panic("failure when reading input data")
	}
}

func solvePuzzleA(input local.Data, result chan int) {
	result <- len(search(input, local.Track{}))
}

func solvePuzzleB(input local.Data, result chan int) {
	result <- 0
}

func search(path *shared.Node, route local.Track) (routes []string) {
	route.Add(path)
	start, end, _, small := route.GetStats()

	if end > 0 {
		routes = append(routes, route.String())
	} else if start < 2 && small < 2 {
		for _, link := range path.Links {
			routes = append(routes, search(link, route)...)
		}
	}

	return
}
