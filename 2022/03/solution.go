package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/shared"
	"sync"
)

func main() {
	input := shared.ReadRuneSlicesFromLines("./input.txt")

	if input != nil {
		answer1 := make(chan int)
		answer2 := make(chan int)

		go solvePart1(input, answer1)
		go solvePart2(input, answer2)

		fmt.Println(<-answer1)
		fmt.Println(<-answer2)
	} else {
		panic("failure when reading input data")
	}
}

func solvePart1(input [][]rune, result chan int) {
	priority := make(chan int, len(input))

	var wg sync.WaitGroup

	for _, entry := range input {
		wg.Add(1)

		go func(entry []rune) {
			defer wg.Done()

			boundary := len(entry) / 2

			compartment1 := entry[:boundary]
			compartment2 := entry[boundary:]

			item := 0

			for _, element := range compartment1 {
				if shared.Contains(compartment2, element) {
					if element < 91 {
						item = int(element) - 38
						break
					}

					if element > 96 {
						item = int(element) - 96
						break
					}
				}
			}

			priority <- item
		}(entry)
	}

	wg.Wait()
	close(priority)

	result <- shared.SumOf(priority)
}

func solvePart2(input [][]rune, result chan int) {
	count := len(input) / 3

	priority := make(chan int, count)

	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)

		begin := i * 3
		end := begin + 3

		go func(entry [][]rune) {
			defer wg.Done()

			item := 0

			for _, element := range entry[0] {
				if shared.Contains(entry[1], element) && shared.Contains(entry[2], element) {
					if element < 91 {
						item = int(element) - 38
						break
					}

					if element > 96 {
						item = int(element) - 96
						break
					}
				}
			}

			priority <- item
		}(input[begin:end])
	}

	wg.Wait()
	close(priority)

	result <- shared.SumOf(priority)
}
