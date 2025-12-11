package pascal

func Triangle(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	pascal := make([][]int, n)
	for i := range pascal {
		pascal[i] = make([]int, i+1)
		pascal[i][0], pascal[i][i] = 1, 1
		for j := 1; j < i; j++ {

			pascal[i][j] = pascal[i-1][j-1] + pascal[i-1][j]

		}

	}
	return pascal
}
