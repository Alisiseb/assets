package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.ErrUnsupported
	}
	hammingValue := 0
	for i, v := range stringToSlice(a) {
		if stringToSlice(b)[i] != v {
			hammingValue++
		}
	}
	return hammingValue, nil
}
func stringToSlice(s string) []rune {
	var sample []rune
	for _, v := range s {
		sample = append(sample, v)
	}
	return sample
}
