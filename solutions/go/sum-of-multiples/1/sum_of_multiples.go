package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var Points = make(map[int]bool)

	sum := 0
	for _, v := range divisors {
		if v == 0 {
			continue
		}
		for i := v; i < limit; i += v {
			if !Points[i] {
				Points[i] = true
			}
		}
	}

	for k, v := range Points {
		if v {
			sum += k
		}
	}
	return sum

}
