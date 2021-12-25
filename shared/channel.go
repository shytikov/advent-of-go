package shared

// GetMaxFromCannel returns largest number found among the cannel contents
func GetMaxFromChannel(input chan int) (result int) {
	for number := range input {
		if result < number {
			result = number
		}
	}

	return
}

// GetSumFromChannel returs sum of all numbers found in the channel
func GetSumFromChannel(input chan int) (result int) {
	for number := range input {
		result += number
	}

	return
}
