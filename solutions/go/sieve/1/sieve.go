package sieve

import "math"

func Sieve(limit int) []int {
	primeSlice := []int{}
	for i := 2; i <= limit; i++ {
		if primeFactors(i) {
			primeSlice = append(primeSlice, i)
		}
	}
	return primeSlice

}
func primeFactors(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		if n%i == 0 {
			return false

		}
	}
	return true
}
