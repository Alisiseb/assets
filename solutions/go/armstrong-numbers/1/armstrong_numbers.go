package armstrong

import "strconv"

func IsNumber(n int) bool {
	nString := strconv.Itoa(n)
	nLen := len(nString)
	var sumArm int
	for _, v := range nString {
		value := int(v - '0')
		sumArm += powInt(value, nLen)
	}
	return n == sumArm
}
func powInt(x, y int) int {
	res := 1
	for i := 0; i < y; i++ {
		res *= x
	}
	return res
}
