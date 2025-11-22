package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
var valueColor = []string{
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

func Value(colors []string) int {
	output := ""
	count := 0
	for _, v := range colors {
		for j, val := range valueColor {
			if val == v {
				count++
				output += strconv.Itoa(j)
			}
		}
		if count == 2 {
			break
		}
	}
	num, _ := strconv.Atoi(output)
	return num
}
