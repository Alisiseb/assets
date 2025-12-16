package wordsearch

import (
	"errors"
	"strings"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	errInPuzzle := validatePuzzle(puzzle)
	if errInPuzzle != nil {
		return nil, errInPuzzle
	}
	isFound := false
	errInFinding := error(nil)
	for _, word := range words {
		for i := range puzzle {
			if strings.Contains(puzzle[i], string(word[0])) {
				for j := 0; j < len(puzzle[0]); j++ {
					if res, ok := leftToRightSearch(word, puzzle, i, j); ok {
						result[word] = res
						isFound = true
						break
					}
					if res, ok := rightToleft(word, puzzle, i, j); ok {
						result[word] = res
						isFound = true
						break
					}
					if res, ok := topToBottom(word, puzzle, i, j); ok {
						result[word] = res
						isFound = true
						break
					}
					if res, ok := bottomToTop(word, puzzle, i, j); ok {
						result[word] = res
						isFound = true
						break
					}
					if res, ok := diagonalTopLeftToBottomRight(word, puzzle, i, j); ok {
						result[word] = res
						isFound = true
						break
					}
					if res, ok := diagonalBottomRightToTopLeft(word, puzzle, i, j); ok {
						result[word] = res
						isFound = true
						break
					}
					if res, ok := diagonalTopRightToBottomLeft(word, puzzle, i, j); ok {
						result[word] = res
						isFound = true
						break
					}
					if res, ok := diagonalBottomLeftToTopRight(word, puzzle, i, j); ok {
						result[word] = res
						isFound = true
						break
					}
				}

			}
		}
		if !isFound {
			result[word] = [2][2]int{{-1, -1}, {-1, -1}}
			errInFinding = errors.New("there is some error in finding words")
		} else {
			isFound = false
		}
	}

	return result, errInFinding
}

func validatePuzzle(puzzle []string) error {
	if len(puzzle) == 0 {
		return errors.New("puzzle cannot be empty")
	}
	rowLength := len(puzzle[0])
	for _, row := range puzzle {
		if len(row) != rowLength {
			return errors.New("all rows in the puzzle must have the same length")
		}
	}
	return nil
}

func leftToRightSearch(word string, puzzle []string, i, j int) ([2][2]int, bool) {
	var result [2][2]int
	wordLength := len(word)
	puzzleWidth := len(puzzle[0])

	if j+wordLength > puzzleWidth {
		return result, false
	}

	for k := 0; k < wordLength; k++ {
		if puzzle[i][j+k] != word[k] {
			return result, false
		}
	}
	result[0] = [2]int{j, i}
	result[1] = [2]int{j + wordLength - 1, i}
	return result, true
}

func rightToleft(word string, puzzle []string, i, j int) ([2][2]int, bool) {
	var result [2][2]int
	wordLength := len(word)
	if j-wordLength+1 < 0 {
		return result, false
	}

	for k := 0; k < wordLength; k++ {
		if puzzle[i][j-k] != word[k] {
			return result, false
		}
	}
	result[0] = [2]int{j, i}
	result[1] = [2]int{j - wordLength + 1, i}
	return result, true
}

func topToBottom(word string, puzzle []string, i, j int) ([2][2]int, bool) {
	var result [2][2]int
	wordLength := len(word)
	puzzleHeight := len(puzzle)

	if i+wordLength > puzzleHeight {
		return result, false
	}

	for k := 0; k < wordLength; k++ {
		if puzzle[i+k][j] != word[k] {
			return result, false
		}
	}
	result[0] = [2]int{j, i}
	result[1] = [2]int{j, i + wordLength - 1}
	return result, true
}

func bottomToTop(word string, puzzle []string, i, j int) ([2][2]int, bool) {
	var result [2][2]int
	wordLength := len(word)
	if i-wordLength+1 < 0 {
		return result, false
	}

	for k := 0; k < wordLength; k++ {
		if puzzle[i-k][j] != word[k] {
			return result, false
		}
	}
	result[0] = [2]int{j, i}
	result[1] = [2]int{j, i - wordLength + 1}
	return result, true
}
func diagonalTopLeftToBottomRight(word string, puzzle []string, i, j int) ([2][2]int, bool) {
	var result [2][2]int
	wordLength := len(word)
	puzzleHeight := len(puzzle)
	puzzleWidth := len(puzzle[0])

	if i+wordLength > puzzleHeight || j+wordLength > puzzleWidth {
		return result, false
	}

	for k := 0; k < wordLength; k++ {
		if puzzle[i+k][j+k] != word[k] {
			return result, false
		}
	}
	result[0] = [2]int{j, i}
	result[1] = [2]int{j + wordLength - 1, i + wordLength - 1}
	return result, true
}

func diagonalBottomRightToTopLeft(word string, puzzle []string, i, j int) ([2][2]int, bool) {
	var result [2][2]int
	wordLength := len(word)

	if i-wordLength+1 < 0 || j-wordLength+1 < 0 {
		return result, false
	}

	for k := 0; k < wordLength; k++ {
		if puzzle[i-k][j-k] != word[k] {
			return result, false
		}
	}
	result[0] = [2]int{j, i}
	result[1] = [2]int{j - wordLength + 1, i - wordLength + 1}
	return result, true
}
func diagonalTopRightToBottomLeft(word string, puzzle []string, i, j int) ([2][2]int, bool) {
	var result [2][2]int
	wordLength := len(word)
	puzzleHeight := len(puzzle)

	if i+wordLength > puzzleHeight || j-wordLength+1 < 0 {
		return result, false
	}

	for k := 0; k < wordLength; k++ {
		if puzzle[i+k][j-k] != word[k] {
			return result, false
		}
	}
	result[0] = [2]int{j, i}
	result[1] = [2]int{j - wordLength + 1, i + wordLength - 1}
	return result, true
}

func diagonalBottomLeftToTopRight(word string, puzzle []string, i, j int) ([2][2]int, bool) {
	var result [2][2]int
	wordLength := len(word)

	if i-wordLength+1 < 0 || j+wordLength > len(puzzle[0]) {
		return result, false
	}

	for k := 0; k < wordLength; k++ {
		if puzzle[i-k][j+k] != word[k] {
			return result, false
		}
	}
	result[0] = [2]int{j, i}
	result[1] = [2]int{j + wordLength - 1, i - wordLength + 1}
	return result, true
}
