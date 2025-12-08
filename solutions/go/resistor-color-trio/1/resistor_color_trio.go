package resistorcolortrio

import (
	"fmt"
	"math"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
var colorValue = []string{
	"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
}

type prefix struct {
	num int
	pre string
}

var metricPrefix = []prefix{
	{9, "giga"},
	{6, "mega"},
	{3, "kilo"},
}

func Label(colors []string) string {
	outputNumber := 0
	prefix := ""
	for i := 0; i < 3; i++ {
		for j, v := range colorValue {
			if v == colors[i] && i < 2 {
				outputNumber += j * int(math.Pow10(1-i))
			} else if v == colors[i] {
				outputNumber *= int(math.Pow10(j))
			}

		}
	}
	for _, v := range metricPrefix {
		if outputNumber/int(math.Pow10(v.num)) > 1 {
			outputNumber /= int(math.Pow10(v.num))
			prefix = v.pre
		}
	}

	return fmt.Sprintf("%d %sohms", outputNumber, prefix)

}
