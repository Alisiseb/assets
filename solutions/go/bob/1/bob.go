// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	responces := []string{"Sure.", "Whoa, chill out!", "Calm down, I know what I'm doing!", "Fine. Be that way!", "Whatever."}
	l := len(remark)
	if l == 0 {
		return responces[3]
	}
	upper, whitespace := Isupper(remark)
	question := questionmark(remark, l)
	switch {
	case question && upper:
		return responces[2]
	case question:
		return responces[0]
	case upper:
		return responces[1]
	case whitespace:
		return responces[3]
	default:
		return responces[4]

	}
}
func questionmark(s string, l int) bool {
	for i := l - 1; i >= 0; i-- {
		if unicode.IsSpace(rune(s[i])) {
			continue
		}
		return rune(s[i]) == '?'
	}
	return false
}

func Isupper(s string) (bool, bool) {
	white := false
	upper := false

	for _, v := range s {
		if unicode.IsUpper(v) {
			upper = true
			continue
		} else if !unicode.IsLetter(v) {
			continue
		} else {
			upper = false
			break
		}

	}
	if !upper {
		for _, v := range s {
			if unicode.IsSpace(v) {
				continue
			} else {
				white = false
				return upper, white
			}

		}
		white = true
	}
	return upper, white
}
