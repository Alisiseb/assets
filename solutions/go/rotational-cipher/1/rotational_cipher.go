package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	output := ""
	for _, v := range plain {
		switch {
		case 'A' <= v && v <= 'Z':
			if int(v)+shiftKey > 'Z' {
				output += string(rune(int(v) + shiftKey - 26))
			} else {
				output += string(rune(int(v) + shiftKey))
			}
		case 'a' <= v && v <= 'z':
			if int(v)+shiftKey > 'z' {
				output += string(rune(int(v) + shiftKey - (26)))
			} else {
				output += string(rune(int(v) + shiftKey))
			}
		default:
			output += string(v)
		}
	}
	return output
}
