package l_e_0118_array

func generate(numRows int) [][]int {
	triangle := make([][]int, 0, numRows)

	// add first
	triangle = append(triangle, make([]int, 1))
	triangle[0][0] = 1

	for row := 1; row < numRows; row++ {
		triangle = append(triangle, make([]int, row+1))

		triangle[row][0], triangle[row][row] = 1, 1
		for i := 1; i <= row/2; i++ {
			triangle[row][i] = triangle[row-1][i-1] + triangle[row-1][i]
			triangle[row][row-i] = triangle[row][i]
		}
	}

	return triangle
}

func Wrapper(numRows int) [][]int {
	return generate(numRows)
}
