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

	if primeMap[n] == 0 {
		primeNumber(n)
	}
	return primeMap[n], nil
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
	if n%2 == 0 {
		return false
	}
	for i := 3; i < n; i = i + 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
