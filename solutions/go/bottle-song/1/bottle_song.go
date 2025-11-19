package bottlesong

import (
	"fmt"
	"strings"
)

var Number = map[int]string{
	0:  "no",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

func firstLine(n int) string {
	if n == 1 {

		return fmt.Sprintf("%v green bottle hanging on the wall,", Number[n])
	}
	return fmt.Sprintf("%v green bottles hanging on the wall,", Number[n])
}
func thirdtLine(n int) string {
	if n == 1 {
		return fmt.Sprintf("There'll be %v green bottle hanging on the wall.", strings.ToLower(Number[n]))
	}
	return fmt.Sprintf("There'll be %v green bottles hanging on the wall.", strings.ToLower(Number[n]))
}

func secLine() string {
	return "And if one green bottle should accidentally fall,"
}

func Recite(startBottles, takeDown int) []string {
	var rec []string

	for i := 0; i < takeDown; i++ {
		if i > 0 {
			rec = append(rec, "")
		}
		if startBottles-i < 0 {
			break
		}
		rec = append(rec, firstLine(startBottles-i), firstLine(startBottles-i), secLine(), thirdtLine(startBottles-i-1))
	}
	return rec

}
