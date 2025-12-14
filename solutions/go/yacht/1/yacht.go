package yacht

import "slices"

func Score(dice []int, category string) int {
	slices.Sort(dice)
	switch category {
	case "ones":
		return sumOfSpecNum(dice, 1)
	case "twos":
		return sumOfSpecNum(dice, 2)
	case "threes":
		return sumOfSpecNum(dice, 3)
	case "fours":
		return sumOfSpecNum(dice, 4)
	case "fives":
		return sumOfSpecNum(dice, 5)
	case "sixes":
		return sumOfSpecNum(dice, 6)
	case "full house":
		countThird := 0
		otherNumbers := make([]int, 0)
		for _, v := range dice {
			if v == dice[2] {
				countThird++
			} else {
				otherNumbers = append(otherNumbers, v)
			}
		}
		if countThird != 3 {
			return 0
		}
		if otherNumbers[0] != otherNumbers[1] {
			return 0
		}
		return (3*dice[2] + 2*otherNumbers[0])
	case "four of a kind":
		count := 0
		for _, v := range dice {
			if v == dice[2] {
				count++
			}
		}
		if count >= 4 {
			return 4 * dice[2]
		}
		return 0
	case "little straight":
		for i, v := range dice {
			if i+1 != v {
				return 0
			}
		}
		return 30
	case "big straight":
		for i, v := range dice {
			if i+2 != v {
				return 0
			}
		}
		return 30
	case "choice":
		sum := 0
		for _, num := range dice {
			sum += num
		}
		return sum
	case "yacht":
		for _, num := range dice {
			if num != dice[0] {
				return 0
			}
		}
		return 50
	default:
		return 0
	}
}
func sumOfSpecNum(dice []int, n int) int {
	sum := 0
	for _, num := range dice {
		if num == n {
			sum += num
		}
	}
	return sum
}
