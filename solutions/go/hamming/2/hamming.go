package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.ErrUnsupported
	}
	hamV := 0
	for i, v := range strings.Split(a, "") {
		if strings.Split(b, "")[i] != v {
			hamV++
		}
	}
	return hamV, nil
}
