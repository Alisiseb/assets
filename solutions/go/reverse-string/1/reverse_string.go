package reverse

func Reverse(input string) string {
	if len(input) <= 1 {
		return input
	}
	result := []rune(input)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
