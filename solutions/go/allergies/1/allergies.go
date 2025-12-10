package allergies

import "math"

var allergenList = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(allergies uint) []string {
	if allergies > 256 {
		allergies = allergies % 256
	}
	var result []string
	for i := len(allergenList) - 1; i >= 0; i-- {
		allergyPoint := uint(math.Pow(2, float64(i)))
		if allergies/(allergyPoint) >= 1 {
			result = append(result, allergenList[i])
			allergies = allergies - allergyPoint
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	allergyitem := Allergies(allergies)
	for _, item := range allergyitem {
		if item == allergen {
			return true
		}
	}
	return false
}
