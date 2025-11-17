package scrabble

import "strings"

func Score(word string) int {
	sc := 0
	for i := 0; i < len(word); i++ {
		for group, letters := range letterG {
			if strings.Contains(letters, strings.ToUpper(word[i:i+1])) {
				sc += letterGScore[group]
			}
		}
	}
	return sc
}

var letterGScore = map[string]int{
	"1g": 1,
	"2g": 2,
	"3g": 3,
	"4g": 4,
	"5g": 5,
	"6g": 8,
	"7g": 10,
}

var letterG = map[string]string{
	"1g": "AEIOULNRST",
	"2g": "DG",
	"3g": "BCMP",
	"4g": "FHVWY",
	"5g": "K",
	"6g": "JX",
	"7g": "QZ",
}
