package local

func Score(theirs, ours int) (result int) {
	//
	i := theirs - 1
	j := ours - 1

	// Matrix contains all possible outcomes:
	// * 0 – we lost
	// * 3 – draw
	// * 6 – we won
	// Indexes are encoded values:
	// * 0 – Rock (measures as 1)
	// * 1 – Paper (measures as 2)
	// * 2 – Scissors (measures as 3)
	solution := [][]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}

	// Round outcome + value of our pick
	return solution[i][j] + ours
}
