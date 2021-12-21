package shared

func GetMaxFromChannel(input chan int) (result int) {
	for number := range input {
		if result < number {
			result = number
		}
	}

	return
}

func GetSumFromChannel(input chan int) (result int) {
	for number := range input {
		result += number
	}

	return
}



