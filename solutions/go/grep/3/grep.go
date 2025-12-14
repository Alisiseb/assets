package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type grepFlags struct {
	l, i, n, v, x bool
}

func Search(pattern string, flags, files []string) []string {
	var oneFile bool
	var flagsOfGrep grepFlags
	totalResults := make([]string, 0)
	if len(flags) != 0 {
		fillFlagMap(flags, &flagsOfGrep)
	}
	if len(files) == 1 {
		oneFile = true
	}
	for _, file := range files {
		results := make([]string, 0)
		rawData, err := os.Open(file)
		if err != nil {
			continue
		}
		results = append(results, grepSearch(rawData, pattern, &flagsOfGrep)...)
		rawData.Close()
		if !oneFile {
			if !flagsOfGrep.l {
				for i := 0; i < len(results); i++ {
					results[i] = fmt.Sprintf("%v:%v", file, results[i])
				}
			}
		}
		totalResults = append(totalResults, results...)
	}

	return totalResults
}

func fillFlagMap(flags []string, gFlags *grepFlags) {
	for _, flag := range flags {
		switch flag {
		case "-l":
			gFlags.l = true
		case "-i":
			gFlags.i = true
		case "-n":
			gFlags.n = true
		case "-v":
			gFlags.v = true
		case "-x":
			gFlags.x = true
		}
	}

}
func isMatched(line, pattern string, exact bool) bool {
	if exact {
		return line == pattern
	}
	return strings.Contains(line, pattern)

}

func grepSearch(file *os.File, pattern string, gFlags *grepFlags) []string {
	var results []string
	scanner := bufio.NewScanner(file)
	if gFlags.i {
		pattern = strings.ToLower(pattern)
	}
	name := file.Name()
	lineNumber := 1
	for scanner.Scan() {
		original := scanner.Text()
		line := original
		if gFlags.i {
			line = strings.ToLower(line)
		}
		match := isMatched(line, pattern, gFlags.x)
		if gFlags.v {
			match = !match
		}
		if match {
			if gFlags.l {
				return append(results, name)
			}
			if gFlags.n {
				results = append(results, fmt.Sprintf("%d:%s", lineNumber, original))
			} else {
				results = append(results, original)
			}
		}
		lineNumber++
	}
	return results
}
