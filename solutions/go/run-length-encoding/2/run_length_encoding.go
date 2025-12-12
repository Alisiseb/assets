package encode

import (
	"regexp"
	"strconv"
	"strings"
)

var group = regexp.MustCompile(`([\d]*?)([\D])`)

func RunLengthEncode(input string) string {

	if input == "" {
		return ""
	}
	var (
		result        strings.Builder
		currentLetter rune
		letterCount   int
	)
	for i, letter := range input {
		if i == 0 {
			currentLetter = letter
			letterCount = 1
			continue
		}
		if letter == currentLetter {
			letterCount++
		} else {
			if letterCount > 1 {
				result.WriteString(strconv.Itoa(letterCount))
			}
			result.WriteRune(currentLetter)
			currentLetter = letter
			letterCount = 1
		}

	}
	if letterCount > 1 {
		result.WriteString(strconv.Itoa(letterCount))
	}
	result.WriteRune(currentLetter)

	return result.String()
}

func RunLengthDecode(input string) string {
	result := ""
	if len(input) == 0 {
		return result
	}
	for _, letterGroup := range group.FindAllStringSubmatch(input, -1) {

		if val, err := strconv.Atoi(letterGroup[1]); err == nil {
			result += strings.Repeat(letterGroup[2], val)
		} else {
			result += letterGroup[2]
		}
	}

	return result
}
