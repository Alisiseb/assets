package atbash

import "strings"

func Atbash(s string) string {
	s = strings.ToLower(s)
	var result strings.Builder
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result.WriteRune('z' - (r - 'a'))
		} else if r >= '0' && r <= '9' {
			result.WriteRune(r)
		}
	}
	stripped := result.String()
	output := ""
	for i := 0; i < len(stripped); i += 5 {
		if i+5 > len(stripped) {
			output += stripped[i:] + " "
			break
		}
		output += stripped[i:i+5] + " "
	}
	return strings.TrimSpace(output)
}
