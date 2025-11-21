package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	test := map[rune]bool{}
	for _, v := range input {
		if v >= 'a' && v <= 'z' {
			_, ok := test[v]
			if !ok {
				test[v] = true
			}
		}

	}
	return len(test) == 26

}
