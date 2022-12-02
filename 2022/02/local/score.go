package local

func Score(theirs, ours int) (result int) {
	i := theirs - 1
	j := ours - 1

	solution := [][]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}

	return solution[i][j] + ours
}
