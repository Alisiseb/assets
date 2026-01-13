package isbn

import (
	"unicode"
)

func IsValidISBN(isbn string) bool {
	var (
		counter int
		sum     int
	)

	for _, v := range isbn {
		if unicode.IsDigit(v) {
			sum += int(v-'0') * (10 - counter)
			counter++
		} else if v == 'X' {
			counter++
			sum += 10
		} else if unicode.IsLetter(v) {
			return false
		}
	}
	return (sum)%11 == 0 && counter == 10
}
