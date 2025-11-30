package prime

import "errors"

var primeMap map[int]int = map[int]int{
	1: 2,
	2: 3,
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("invalid input")
	}
	val := primeMap[n]
	for val == 0 {
		primeNumber(n)
		val = primeMap[n]
	}
	return val, nil
}

func primeNumber(n int) {

	for i := 3; i <= n; i++ {
		j := 1
		for {
			newnum := primeMap[i-1] + j
			if prime(newnum) {
				primeMap[i] = newnum
				break
			}
			j++
		}
	}
}

func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
