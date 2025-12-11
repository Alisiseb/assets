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
	for i := range diamond {
		diamond[i] = make([]string, size)
		letter := byte('A' + i)
		diamond[i][diff+i] = string(rune(letter))
		diamond[i][diff-i] = string(rune(letter))
		makeSpace(&diamond[i])
		if letter == char {
			break
		}
	}
	j := 1
	for i := diff + 1; i < size; i++ {
		diamond[i] = make([]string, size)

		letter := rune(int(rune(char)) - j)
		diamond[i][j] = string(letter)
		diamond[i][size-j-1] = string(letter)
		makeSpace(&diamond[i])
		j++
		if j > diff {
			break
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
