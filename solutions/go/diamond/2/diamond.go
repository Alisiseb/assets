package diamond

import (
	"errors"
)

func Gen(char byte) (string, error) {
	diff := int(rune(char) - 'A')
	if diff < 0 || diff > 25 {
		return "", errors.New("input invalid")
	}
	if diff == 0 {
		return "A", nil
	}
	size := diff*2 + 1
	result := ""
	diamond := make([][]string, size)
	//first half
	j := 1
	for i := range diamond {
		diamond[i] = make([]string, size)
		if i <= diff {
			letter := byte('A' + i)
			diamond[i][diff+i] = string(letter)
			diamond[i][diff-i] = string(letter)
			makeSpace(&diamond[i])
		} else {

			letter := (char) - byte(j)
			diamond[i][j] = string(letter)
			diamond[i][size-j-1] = string(letter)
			makeSpace(&diamond[i])
			j++
		}

	}
	for i := range diamond {
		for _, v := range diamond[i] {
			result += v
		}
		if i < size-1 {
			result += "\n"
		}
	}
	return result, nil
}

func makeSpace(s *[]string) {
	for i, v := range *s {
		if v == "" {
			(*s)[i] = " "
		}
	}
}
