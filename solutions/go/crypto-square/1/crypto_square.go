package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	re := regexp.MustCompile(`[A-Za-z0-9]+`)
	slice := re.FindAllString(pt, -1)
	new := ""
	for _, v := range slice {
		new += strings.ToLower(v)
	}
	l := len(new)
	r := int(math.Round(math.Sqrt(float64(l))))
	c := 0
	if r*r >= l {
		c = r
	} else {
		c = r + 1
	}
	result := ""
	groupCounter := 0
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			if i+j*c >= l {
				result += " "
				continue
			}
			result += string(new[i+j*c])
		}
		// if r+i*c <= l && i+1 != c {
		groupCounter++
		if groupCounter < c {

			result += " "
		}

	}
	return result
}
