package shared

import (
	"math"
)

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
