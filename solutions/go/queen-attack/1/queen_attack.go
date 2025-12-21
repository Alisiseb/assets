package queenattack

import (
	"errors"
	"math"
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if valid, err := validatePosition(whitePosition); !valid {
		return false, err
	}
	if valid, err := validatePosition(blackPosition); !valid {
		return false, err
	}
	if whitePosition == blackPosition {
		return false, errors.New("queens cannot occupy the same position")
	}
	whitePos := []rune(whitePosition)
	blackPos := []rune(blackPosition)
	if blackPos[0] == whitePos[0] || whitePos[1] == blackPos[1] {
		return true, nil
	}
	if math.Abs(float64(whitePos[0]-blackPos[0])) == math.Abs(float64(whitePos[1]-blackPos[1])) {
		return true, nil
	} else {
		return false, nil
	}

}
func validatePosition(position string) (bool, error) {
	if len(position) != 2 {
		return false, errors.New("invalid position")
	}
	col := position[0]
	row := position[1]
	if col < 'a' || col > 'h' {
		return false, errors.New("invalid position column")
	}
	if row < '1' || row > '8' {
		return false, errors.New("invalid position row")
	}
	return true, nil
}
