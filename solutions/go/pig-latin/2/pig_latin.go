package piglatin

import (
	"regexp"
	"strings"
)

var (
	rule1 = regexp.MustCompile(`^(a|e|i|o|u|xr|yt)`)
	rule2 = regexp.MustCompile(`^([^aeiou]+)\w*`)
	rule3 = regexp.MustCompile(`^([^aeiouy]*qu)\w*`)
	rule4 = regexp.MustCompile(`^([^aeiouy]+)(y\w*)`)
)

func Sentence(sentence string) string {
	re := regexp.MustCompile(`\s`)
	words := re.Split(sentence, -1)
	for i, word := range words {
		if newWord, ok := ruleOne(word); ok {
			words[i] = newWord
			continue
		}
		if newWord, ok := ruleThree(word); ok {
			words[i] = newWord
			continue
		}
		if newWord, ok := ruleFour(word); ok {
			words[i] = newWord
			continue
		}
		if newWord, ok := ruleTwo(word); ok {
			words[i] = newWord
			continue
		}
	}

	return strings.Join(words, " ")

}

func ruleOne(word string) (string, bool) {

	if rule1.MatchString(word) {
		return word + "ay", true
	}
	return "", false
}

func ruleTwo(word string) (string, bool) {
	submatches := rule2.FindStringSubmatch(word)
	if submatches == nil {
		return "", false
	}

	return submatches[2] + submatches[1] + "ay", true

}

func ruleThree(word string) (string, bool) {
	submatches := rule3.FindStringSubmatch(word)
	if submatches == nil {
		return "", false
	}
	return submatches[2] + submatches[1] + "ay", true
}

func ruleFour(word string) (string, bool) {

	submatches := rule4.FindStringSubmatch(word)
	if submatches == nil {

		return "", false
	}

	//prefix := re.FindString(word)
	return submatches[2] + submatches[1] + "ay", true

}
