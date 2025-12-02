package series

func All(n int, s string) []string {
	if n > len(s) || n <= 0 {
		return nil
	}
	result := []string{}
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) || n <= 0 {
		return ""
	}
	return s[0:n]
}
