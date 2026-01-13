package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	output := []string{}
	subject = strings.ToLower(subject)
	subjectLetterMap := letterCount(subject)
	for j, v := range candidates {
		v = strings.ToLower(v)
		if subject == v || len(v) != len(subject) {
			continue
		}
		checker := letterCheck(v, subject)

		candidateLetterMap := letterCount(v)
		if checker && letterMapCheck(subjectLetterMap, candidateLetterMap) {

			output = append(output, candidates[j])

		}

	}
	return output
}
func letterCount(s string) map[rune]int {
	letterMap := make(map[rune]int)
	for _, l := range s {
		letterMap[l] = letterMap[l] + 1
	}
	return letterMap
}

func letterMapCheck(m1, m2 map[rune]int) bool {
	for k1, v1 := range m1 {
		if m2[k1] != v1 {
			return false
		}
	}
	return true
}
func letterCheck(v, subject string) bool {
	checker := false
	for _, r := range v {
		checker = false
		for _, rr := range subject {
			if r == rr {
				checker = true

				break
			}
		}
		if checker {
			continue
		} else {
			break
		}
	}
	return checker
}
