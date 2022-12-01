package shared

// GetMaxFromChannel returns the largest number found among the channel contents
func GetMaxFromChannel(input chan int) (result int) {
	for number := range input {
		if result < number {
			result = number
		}
	}

	return
}

// GetSumFromChannel returns sum of all numbers found in the channel
func GetSumFromChannel(input chan int) (result int) {
	for number := range input {
		result += number
	}

	return
}
