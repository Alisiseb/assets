package piglatin

import (
	"regexp"
	"strings"
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
	vowelOneRe := regexp.MustCompile(`^(a|e|i|o|u|xr|yt)`)
	if vowelOneRe.MatchString(word) {
		return word + "ay", true
	}
	return "", false
}

func ruleTwo(word string) (string, bool) {
	re := regexp.MustCompile(`^([^aeiou]+)`)
	if re.MatchString(word) {
		prefix := re.FindString(word)
		return word[len(prefix):] + prefix + "ay", true
	}
	return "", false

}

func ruleThree(word string) (string, bool) {
	re := regexp.MustCompile(`^([^aeiouy]*qu)`)
	if re.MatchString(word) {
		prefix := re.FindString(word)
		return word[len(prefix):] + prefix + "ay", true
	}
	return "", false
}

func ruleFour(word string) (string, bool) {
	re := regexp.MustCompile(`^([^aeiouy]+)(y\w*)`)
	submatches := re.FindStringSubmatch(word)
	if submatches == nil {

		return "", false
	}

	//prefix := re.FindString(word)
	return submatches[2] + submatches[1] + "ay", true

}
