package pascal

func Triangle(n int) [][]int {
	pascal := make([][]int, n)
	for i := range pascal {
		pascal[i] = make([]int, i+1)
		if i == 0 {
			pascal[i][0] = 1
			continue
		}
		for j := 0; j <= i; j++ {
			upperLeftValue := 0
			upperRightValue := 0
			if j-1 >= 0 {
				upperLeftValue = pascal[i-1][j-1]
			}
			if j <= i-1 {
				upperRightValue = pascal[i-1][j]
			}
			pascal[i][j] = upperLeftValue + upperRightValue
		}

	}
	return pascal
}
