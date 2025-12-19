package wordsearch

import (
	"errors"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	errInPuzzle := validatePuzzle(puzzle)
	directions := []struct {
		dx, dy int
	}{
		{1, 0},   // →
		{-1, 0},  // ←
		{0, 1},   // ↓
		{0, -1},  // ↑
		{1, 1},   // ↘
		{-1, -1}, // ↖
		{-1, 1},  // ↙
		{1, -1},  // ↗
	}
	if errInPuzzle != nil {
		return nil, errInPuzzle
	}
	errInFinding := error(nil)
	for _, word := range words {
		for i := range puzzle {
			for j := 0; j < len(puzzle[0]); j++ {
				for _, dir := range directions {
					if res, ok := searchForWordBaseOnFirstcharacter(word, puzzle, i, j, dir.dx, dir.dy); ok {
						result[word] = res
						goto NextWord
					}
				}
			}
		}
		result[word] = [2][2]int{{-1, -1}, {-1, -1}}
		errInFinding = errors.New("there is some error in finding words")
	NextWord:
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

func searchForWordBaseOnFirstcharacter(word string, puzzle []string, i, j, dx, dy int) ([2][2]int, bool) {
	var result [2][2]int
	wordLength := len(word)
	rowSize := len(puzzle[0])
	colSize := len(puzzle)

	if !validateDirection(i, j, wordLength, rowSize, colSize, dx, dy) {
		return result, false
	}

	for k := range word {
		if puzzle[i+dy*k][j+dx*k] != word[k] {
			return result, false
		}
	}
	result[0] = [2]int{j, i}
	result[1] = [2]int{j + dx*(wordLength-1), i + dy*(wordLength-1)}
	return result, true
}

func validateDirection(i, j, wordLength, maxRows, maxCols, dx, dy int) bool {
	newI := i + dy*(wordLength-1)
	newJ := j + dx*(wordLength-1)

	if newI < 0 || newI >= maxCols || newJ < 0 || newJ >= maxRows {
		return false
	}
	return true
}
