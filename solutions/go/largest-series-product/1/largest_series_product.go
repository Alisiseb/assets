package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	largestSeies := 0
	if span > len(digits) || span < 0 {
		return int64(largestSeies), errors.New("invalid input")
	}
	for _, v := range digits {
		if !unicode.IsDigit(v) {
			return int64(largestSeies), errors.New("invalid input")
		}
	}
	for i := 0; i < len(digits); i++ {
		if i+span > len(digits) {
			break
		}
		if largestSeies < multiply(digits[i:i+span]) {

			largestSeies = multiply(digits[i : i+span])
		}

	}
	return int64(largestSeies), nil
}

func multiply(dig string) int {
	sum := 1
	for i := 0; i < len(dig); i++ {
		digit, _ := strconv.Atoi(string(dig[i]))
		sum *= digit
	}
	return (sum)
}
