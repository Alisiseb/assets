package cipher

import "strings"

// Define the shift and vigenere types here.
type shift struct {
	distance int
}

type vigenere struct {
	key string
}

// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	return shift{distance: distance}
}

func (c shift) Encode(input string) string {
	input = strings.ToLower(input)
	output := ""
	if c.distance < 0 {
		c.distance = (c.distance%26 + 26) % 26
	}
	for _, char := range input {

		if char >= 'a' && char <= 'z' {
			shifted := ((int(char-'a') + c.distance) % 26) + 'a'
			output += string(rune(shifted))
		}
	}
	return output
}

func (c shift) Decode(input string) string {
	input = strings.ToLower(input)
	output := ""
	if c.distance < 0 {
		c.distance = (c.distance%26 + 26) % 26
	}
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			shifted := ((int(char-'a') - c.distance + 26) % 26) + 'a'
			output += string(rune(shifted))
		}
	}
	return output

}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	no_a_char := true
	for _, char := range key {
		if char < 'a' || char > 'z' {
			return nil
		}
		if char != 'a' {
			no_a_char = false
		}
	}
	if no_a_char {
		return nil
	}
	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	output := ""
	skipchar := 0
	for i, char := range input {
		if char >= 'a' && char <= 'z' {
			shiftchar := shift{distance: int(v.key[(i-skipchar)%len(v.key)] - 'a')}
			output += shiftchar.Encode(string(char))
		} else {
			skipchar++
		}

	}
	return output
}

func (v vigenere) Decode(input string) string {
	input = strings.ToLower(input)
	output := ""
	skipchar := 0
	for i, char := range input {
		if char >= 'a' && char <= 'z' {
			shiftchar := shift{distance: int(v.key[(i-skipchar)%len(v.key)] - 'a')}
			output += shiftchar.Decode(string(char))
		} else {
			skipchar++
		}
	}
	return output
}
