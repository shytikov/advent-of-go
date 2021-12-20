package utils

import "math"

func SumInts(slice []int) (result int) {
	for _, value := range slice {
		result += value
	}

	return
}

func MaxInInts(slice []int) (result int) {
	for _, value := range slice {
		if result < value {
			result = value
		}
	}

	return
}

// Converts loose binaries – array of integer `0` and `1` as a single integer
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
