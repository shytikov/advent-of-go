package utils

func SumNumbers(slice []int) int {
	result := 0

	for _, value := range slice {
		result += value
	}

	return result
}
