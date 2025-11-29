package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	wordC := make(Frequency)
	phrase = strings.ToLower(phrase)
	re := regexp.MustCompile(`[!@#$%^&*,.:"]+`)

	for re.MatchString(phrase) {

		phrase = re.ReplaceAllLiteralString(phrase, " ")
	}
	wordSlice := strings.Fields(phrase)
	re = regexp.MustCompile(`^'|'$`)
	for _, v := range wordSlice {
		v = re.ReplaceAllString(v, "")
		if v == "" {
			continue
		}
		wordC[v] = wordC[v] + 1
	}
	return wordC
}
