package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	s := strings.ReplaceAll(question, "?", "")
	s = strings.ReplaceAll(s, "What is ", "")
	l := len(s)
	nc, nums, numIndex := numCount(s)
	oc, ops, opsIndex := opsCount(s)
	if !validation(nc, oc, l, numIndex, opsIndex) {
		return 0, false
	}
	switch nc {
	case 1:
		return nums[0], true
	case 2:
		return operation(ops[0], nums[0], nums[1])
	case 3:
		newNum, check := operation(ops[0], nums[0], nums[1])
		if !check {
			return 0, check
		}
		return operation(ops[1], newNum, nums[2])
	default:
		return 0, false
	}
}

func numCount(s string) (int, []int, [][]int) {
	nums := []int{}
	re := regexp.MustCompile(`[-]?\d+`)
	count := re.FindAllString(s, -1)
	numIndex := re.FindAllStringIndex(s, -1)

	for _, v := range count {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}
	return len(count), nums, numIndex
}

func opsCount(s string) (int, []string, [][]int) {
	re := regexp.MustCompile(`(?:plus|divided by|multiplied by|minus)`)
	count := re.FindAllString(s, -1)
	opIndex := re.FindAllStringIndex(s, -1)

	return len(count), count, opIndex
}
func validation(n, m, l int, ni, oi [][]int) bool {
	if n <= m {
		return false
	}
	if n != m+1 {
		return false
	}
	switch n {
	case 1:
		if ni[0][1] != l {
			return false
		}
	default:
		if ni[0][1] > oi[0][0] || ni[1][0] < oi[0][0] {
			return false
		}
	}

	return true
}
func operation(ops string, n, m int) (int, bool) {
	switch ops {
	case "plus":
		return n + m, true
	case "minus":
		return n - m, true
	case "divided by":
		if m == 0 {
			return 0, false
		}
		return n / m, true
	case "multiplied by":
		return n * m, true
	default:
		return 0, false
	}
}
