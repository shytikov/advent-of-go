package shared

import (
	"math"
)

// SumOf returns sum of all numbers in the slice
func SumOf(slice []int) (result int) {
	for _, value := range slice {
		result += value
	}

	return
}

// ProductOf returns product of multipliication of all numbers in the slice
func ProductOf(slice []int) (result int) {
	result = 1

	for _, value := range slice {
		result *= value
	}

	return
}

func MaxOf(slice []int) (result int) {
	result = math.MinInt

	for _, value := range slice {
		if result < value {
			result = value
		}
	}

	return
}

func MinOf(slice []int) (result int) {
	result = math.MaxInt

	for _, value := range slice {
		if result > value {
			result = value
		}
	}

	return
}

// LooseBinaryToInt converts loose binaries – array of integer `0` and `1` as a single integer
// For example, 001010011010 will become 666
func LooseBinaryToInt(slice []int) (result int) {
	count := len(slice)

	for i, digit := range slice {
		// Only `0` and `1` are valid at this point
		if digit == 0 || digit == 1 {
			result += int(math.Pow(2, float64(count-i-1))) * digit
		}
	}

	return
}
