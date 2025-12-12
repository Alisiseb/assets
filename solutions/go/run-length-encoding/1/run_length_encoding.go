package encode

import (
	"regexp"
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {

	result := ""
	if input == "" {
		return result
	}
	letters := []rune(input)
	sum := 0
	for i := 0; i < len(letters); i++ {
		letterCount := 0
		letter := letters[i]
		for j := i; j < len(letters); j++ {
			if letter == letters[j] {
				letterCount++

			} else {
				i = j - 1
				break
			}
		}
		sum += letterCount
		if letterCount != 1 {
			result += strconv.Itoa(letterCount)
		}
		result += string(letter)
		if sum == len(letters) {
			break
		}
	}

	return result
}

func RunLengthDecode(input string) string {
	result := ""
	if len(input) == 0 {
		return result
	}
	group := regexp.MustCompile(`([\d]*?)([\D])`)
	letterCombo := group.FindAllString(input, -1)
	for i := 0; i < len(letterCombo); i++ {
		lettergroup := group.FindStringSubmatch(letterCombo[i])
		if val, err := strconv.Atoi(lettergroup[1]); err == nil {
			result += strings.Repeat(lettergroup[2], val)
		} else {
			result += lettergroup[2]
		}
	}

	return result
}
