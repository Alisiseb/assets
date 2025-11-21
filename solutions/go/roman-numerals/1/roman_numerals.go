package romannumerals

import "errors"

var RomanNumMap = []struct {
	romanNumber string
	num         int
}{

	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("invalid input")
	}
	return RomanWriter(input), nil
}
func RomanWriter(n int) string {
	if n == 0 {
		return ""
	}
	result := ""
	for _, v := range RomanNumMap {
		for n >= v.num {
			result += v.romanNumber
			n = n - v.num
		}

	}
	return result
}
