package local

func Score(theirs, ours rune) (result int) {
	i := int(theirs) - 64
	j := int(ours) - 87

	matrix := [][]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}

	return matrix[i-1][j-1] + j
}
