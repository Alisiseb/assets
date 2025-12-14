package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// - `-i` Match using a case-insensitive comparison.

// - `-l` Output only the names of the files that contain at least one matching line.
// - `-n` Prepend the line number and a colon (':') to each line in the output, placing the number after the filename (if present).

// - `-v` Invert the program -- collect all lines that fail to match.
// - `-x` Search only for lines where the search string matches the entire line.
var flagsMap = make(map[string]bool, 5)

func Search(pattern string, flags, files []string) []string {
	var (
		oneFile bool
	)
	totalResults := make([]string, 0)

	flagsMap = map[string]bool{
		"-l": false,
		"-i": false,
		"-n": false,
		"-v": false,
		"-x": false,
	}
	if len(flags) != 0 {
		fillFlagMap(flags)
	}
	if len(files) == 1 {
		oneFile = true
	}
	for _, file := range files {
		results := make([]string, 0)
		rawData, err := os.Open(file)
		if err != nil {
		}
		defer rawData.Close()

		if flagsMap["-l"] {
			if flagLIsCalled(rawData, pattern, flagsMap["-i"], flagsMap["-n"], flagsMap["-v"]) {
				results = append(results, file)
				totalResults = append(totalResults, results...)
				continue
			}
		}

		if flagsMap["-x"] {
			results = append(results, flagXIsCalled(rawData, pattern, flagsMap["-i"], flagsMap["-n"], flagsMap["-v"])...)
		} else {
			results = append(results, flagIsCalled(rawData, pattern, flagsMap["-i"], flagsMap["-n"], flagsMap["-v"])...)
		}
		if !oneFile {
			for i := 0; i < len(results); i++ {
				results[i] = fmt.Sprintf("%v:%v", file, results[i])
			}
		}
		totalResults = append(totalResults, results...)
	}

	return totalResults
}

func fillFlagMap(flags []string) {
	for _, flag := range flags {
		for key := range flagsMap {
			if flag == key {
				flagsMap[key] = true
				break
			}
		}
	}

}
func isMatched(line, pattern string) bool {
	return strings.Contains(line, pattern)
}

func flagLIsCalled(file *os.File, pattern string, caseinsensitive bool, flagNpresented bool, flagVpresented bool) bool {
	result := flagIsCalled(file, pattern, caseinsensitive, flagNpresented, flagVpresented)
	return len(result) > 0
}

func flagIsCalled(file *os.File, pattern string, caseinsensitive bool, flagNpresented bool, flagVpresented bool) []string {
	var results []string
	scanner := bufio.NewScanner(file)
	if caseinsensitive {
		pattern = strings.ToLower(pattern)
	}
	lineNumber := 1
	for scanner.Scan() {
		original := scanner.Text()
		line := original
		if caseinsensitive {
			line = strings.ToLower(line)
		}
		if flagVpresented {
			if !isMatched(line, pattern) {
				if flagNpresented {
					results = append(results, fmt.Sprintf("%d:%s", lineNumber, original))
				} else {
					results = append(results, original)
				}
			}
		} else {
			if isMatched(line, pattern) {
				if flagNpresented {
					results = append(results, fmt.Sprintf("%d:%s", lineNumber, original))
				} else {
					results = append(results, original)
				}
			}
		}
		lineNumber++
	}
	return results
}

func flagXIsCalled(file *os.File, pattern string, caseinsensitive bool, flagNpresented bool, flagVpresented bool) []string {
	var results []string
	scanner := bufio.NewScanner(file)
	if caseinsensitive {
		pattern = strings.ToLower(pattern)
	}
	lineNumber := 1
	for scanner.Scan() {
		original := scanner.Text()
		line := original
		if caseinsensitive {
			line = strings.ToLower(line)
		}
		if flagVpresented {
			if line != pattern {
				if flagNpresented {
					results = append(results, fmt.Sprintf("%d:%s", lineNumber, original))
				} else {
					results = append(results, original)
				}

			}
		} else {
			if line == pattern {
				if flagNpresented {
					results = append(results, fmt.Sprintf("%d:%s", lineNumber, original))
				} else {
					results = append(results, original)
				}

			}
		}

		lineNumber++
	}
	return results
}
