package main

import (
	"fmt"
	"strings"

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
	criterion := func(route local.Track) (startFlag, endFlag, doneFlag bool) {
		counts := map[string]int{
			"start": 0,
			"end":   0,
			"small": 0,
		}

		for name, count := range route.Visited {
			if name == "start" {
				counts["start"] += count
			} else if name == "end" {
				counts["end"] += count
			} else {
				if strings.ToLower(name) == name {
					if counts["small"] < count {
						counts["small"] = count
					}
				}
			}
		}

		return counts["start"] > 1, counts["end"] > 0, counts["small"] > 1
	}

	result <- len(search(input, local.Track{}, criterion))
}

func solvePuzzleB(input local.Data, result chan int) {
	result <- 0
}

func search(cave *shared.Node, route local.Track, resolver local.Criterion) (routes []string) {
	route.Add(cave)

	start, end, done := resolver(route)

	if end {
		routes = append(routes, route.String())
	} else if !start && !done {
		for _, link := range cave.Links {
			routes = append(routes, search(link, route, resolver)...)
		}
	}

	return
}
