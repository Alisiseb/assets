package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		results = append(results, flagIsCalled(rawData, pattern)...)
		defer rawData.Close()
		if !oneFile {
			if !flagsMap["-l"] {
				for i := 0; i < len(results); i++ {
					results[i] = fmt.Sprintf("%v:%v", file, results[i])
				}
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
	if flagsMap["-x"] {
		return line == pattern
	}
	return strings.Contains(line, pattern)

}

func flagIsCalled(file *os.File, pattern string) []string {
	var results []string
	scanner := bufio.NewScanner(file)
	if flagsMap["-i"] {
		pattern = strings.ToLower(pattern)
	}
	name := file.Name()
	lineNumber := 1
	for scanner.Scan() {
		original := scanner.Text()
		line := original
		if flagsMap["-i"] {
			line = strings.ToLower(line)
		}
		if flagsMap["-v"] {
			if !isMatched(line, pattern) {
				if flagsMap["-l"] {
					return append(results, name)
				}
				if flagsMap["-n"] {
					results = append(results, fmt.Sprintf("%d:%s", lineNumber, original))
				} else {
					results = append(results, original)
				}
			}
		} else {
			if isMatched(line, pattern) {
				if flagsMap["-l"] {
					return append(results, name)
				}
				if flagsMap["-n"] {
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
