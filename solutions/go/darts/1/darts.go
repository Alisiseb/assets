package darts

import "math"

func Score(x, y float64) int {
	dis := math.Pow(x*x+y*y, 0.5)
	if dis > 10 {
		return 0
	} else if dis > 5 {
		return 1
	} else if dis > 1 {
		return 5
	} else {
		return 10
	}
}
