package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid input")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	sum := 0
	for i := 0; i < 64; i++ {
		num, _ := Square(i + 1)
		sum += int(num)

	}
	return uint64(sum)
}
