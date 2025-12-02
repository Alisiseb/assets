package prime

func Factors(n int64) []int64 {
	results := []int64{}
	if int(n) <= 1 {
		return results
	}
	for i := 2; i <= int(n); i++ {
		for {
			if int(n)%i == 0 {
				results = append(results, int64(i))
				n = n / int64(i)
			} else {
				break
			}
		}
	}
	return results
}
