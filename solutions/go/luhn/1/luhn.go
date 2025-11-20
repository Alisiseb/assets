package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	l := len(id)
	var idSlice []int
	if l <= 1 {
		return false
	}
	for _, v := range id {
		if unicode.IsDigit(v) {
			num, _ := strconv.Atoi(string(v))
			idSlice = append(idSlice, num)
		} else {
			return false
		}
	}

	for i := 0; i < l; i += 2 {
		if l-i-2 < 0 {
			break
		}
		if idSlice[l-i-2]*2 > 9 {
			idSlice[l-i-2] = idSlice[l-i-2]*2 - 9
		} else {
			idSlice[l-i-2] = idSlice[l-i-2] * 2
		}

	}
	sum := 0
	for _, v := range idSlice {
		sum += v
	}
	return sum%10 == 0
}
