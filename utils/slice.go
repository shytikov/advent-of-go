package utils

import "math"

func SumInts(slice []int) int {
	result := 0

	for _, value := range slice {
		result += value
	}

	return result
}

func BinaryIntsToInt(slice []int) int {
	result := 0
	count := len(slice)

	for i, value := range slice {
		result += int(math.Pow(2, float64(count-i-1))) * value
	}

	return result
}
