package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	err := validateDigits(inputBase, outputBase, inputDigits)
	if err != nil {
		return nil, err
	}
	lengthInput := len(inputDigits)
	if lengthInput == 0 {
		return []int{0}, nil
	}
	numberBaseTen := 0

	for i, v := range inputDigits {
		numberBaseTen += v * int(math.Pow(float64(inputBase), float64(lengthInput-i-1)))
	}

	return ToBase(numberBaseTen, outputBase), nil

}
func validateDigits(iB, oB int, iD []int) error {

	if iB < 2 {
		return errors.New("input base must be >= 2")
	}

	if oB < 2 {
		return errors.New("output base must be >= 2")
	}
	for _, v := range iD {
		if v >= iB || v < 0 {
			return errors.New("all digits must satisfy 0 <= d < input base")
		}
	}
	return nil
}
func ToBase(n, b int) []int {
	outputDigit := make([]int, 0)

	if int(n/b) == 0 {
		outputDigit = append(outputDigit, n)
		return outputDigit
	} else {
		outputDigit = append(outputDigit, ToBase(int(n/b), b)...)
		outputDigit = append(outputDigit, n%b)

	}
	return outputDigit
}
