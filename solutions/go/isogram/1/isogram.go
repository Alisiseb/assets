package isogram

import "strings"

func IsIsogram(word string) bool {
	letterMap := map[rune]bool{}
	for _, v := range strings.ToLower(word) {
		if v >= 'a' && v <= 'z' {

			if !letterMap[v] {
				letterMap[v] = true
			} else {
				return false
			}
		}

	}
	return true
}
