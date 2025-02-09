package l_e_0119_array

func getRow(rowIndex int) []int {
	rowIndex++

	row := make([]int, rowIndex)

	// First and last elements set to 1
	row[0], row[rowIndex-1] = 1, 1

	half := rowIndex / 2
	for i := 1; i < half; i++ {
		row[i] = row[i-1] * (rowIndex - i) / i
		row[rowIndex-1-i] = row[i]
	}

	if rowIndex > 2 && rowIndex%2 == 1 {
		row[half] = row[half-1] * (rowIndex - half) / half
	}

	return row
}

func Wrapper(rowIndex int) []int {
	return getRow(rowIndex)
}
