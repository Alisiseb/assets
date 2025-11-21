package strain

// Implement the "Keep" and "Discard" function in this file.

func Keep[T any](input []T, predicate func(T) bool) []T {
	result := []T{}
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Discard[T any](input []T, predicate func(T) bool) []T {
	result := []T{}
	for _, v := range input {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
