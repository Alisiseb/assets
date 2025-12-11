package perfect

import "errors"

// Define the Classification type here.
type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationAbundant
	ClassificationPerfect
)

var ErrOnlyPositive = errors.New("only positve numbers accepted")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	if n == 1 {
		return ClassificationDeficient, nil
	}

	sum := int64(aliquotsum(int(n)))
	switch {
	case sum > n:
		return ClassificationAbundant, nil
	case sum < n:
		return ClassificationDeficient, nil
	default:
		return ClassificationPerfect, nil
	}
}

func aliquotsum(n int) int {
	sum := 1
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
