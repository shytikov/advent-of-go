package shared

import (
	"log"
)

func ActOn(err error) {
	if err != nil {
		log.Fatal(err) // or panic or os.Exit
	}
}

// MaxOf returns the largest number found among the channel / slice contents
func MaxOf[T chan int | []int](input T) (result int) {
	for number := range input {
		if result < number {
			result = number
		}
	}

	return
}

// MinOf returns the smallest number found among the channel / slice contents
func MinOf[T chan int | []int](input T) (result int) {
	for _, value := range input {
		if result > value {
			result = value
		}
	}

	return
}

// SumOf returns sum of all numbers found in the channel / slice
func SumOf[T chan int | []int](input T) (result int) {
	for number := range input {
		result += number
	}

	return
}

// ProductOf returns product of multiplication of all numbers in the channel / slice
func ProductOf[T chan int | []int](input T) (result int) {
	result = 1

	for _, value := range input {
		result *= value
	}

	return
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
