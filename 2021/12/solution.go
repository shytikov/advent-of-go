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
	criterion := func(route local.Track) (circular, complete, invalid bool) {
		start, end, small := 0, 0, 0

		for name, count := range route.Visited {
			if name == "start" {
				start += count
				circular = start > 1
				if circular {
					break
				}
			} else if name == "end" {
				end += count
				complete = end > 0
				if complete {
					break
				}
			} else if strings.ToLower(name) == name {
				if small < count {
					small = count
					invalid = small > 1
					if invalid {
						break
					}
				}
			}
		}

		return
	}

	result <- len(search(input, local.Track{}, criterion))
}

func solvePuzzleB(input local.Data, result chan int) {
	criterion := func(route local.Track) (circular, complete, invalid bool) {
		start, end, double := 0, 0, 0

		for name, count := range route.Visited {
			if name == "start" {
				start += count
				circular = start > 1
				if circular {
					break
				}
			} else if name == "end" {
				end += count
				complete = end > 0
				if complete {
					break
				}
			} else if strings.ToLower(name) == name {
				if count == 2 {
					double = double + 1
					invalid = double > 1
					if invalid {
						break
					}
				} else if count > 2 {
					invalid = true
					circular = true
					break
				}
			}
		}

		return
	}

	result <- len(search(input, local.Track{}, criterion))
}

func search(cave *shared.Node, route local.Track, resolver local.Criterion) (routes []string) {
	route.Add(cave)

	circular, complete, invalid := resolver(route)

	if complete {
		routes = append(routes, route.String())
	} else if !circular && !invalid {
		for _, link := range cave.Links {
			routes = append(routes, search(link, route, resolver)...)
		}
	}

	return
}
