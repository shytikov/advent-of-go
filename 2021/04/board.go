package main

type Board struct {
	Numbers [dimension][dimension]int
	Rows    [dimension]int
	Columns [dimension]int
	Sum     int
}

func (b Board) draw(number int) (result Board) {
	result = b

	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			if result.Numbers[i][j] == number {
				result.Rows[i] += 1
				result.Columns[j] += 1
				result.Sum = result.Sum - number
			}
		}
	}

	return
}

func (b Board) hasWon() bool {
	for i := 0; i < dimension; i++ {
		if b.Rows[i] == dimension {
			return true
		}
	}

	for i := 0; i < dimension; i++ {
		if b.Columns[i] == dimension {
			return true
		}
	}

	return false
}

func (b Board) getScore(number int) int {
	return b.Sum * number
}
