package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	l := len(id)
	var idSlice []int
	if l <= 1 {
		return false
	}
	for _, v := range id {

		num, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		idSlice = append(idSlice, num)

	}
	sum := 0
	d := 0
	for i := l - 1; i >= 0; i-- {
		if d%2 != 0 {
			if idSlice[i]*2 > 9 {
				sum += idSlice[i]*2 - 9
			} else {
				sum += idSlice[i] * 2
			}
			d++
			continue
		}
		d++
		sum += idSlice[i]

	}
	return sum%10 == 0
}
