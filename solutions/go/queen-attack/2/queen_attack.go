package queenattack

import (
	"errors"
	"math"
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	whitePos := []rune(whitePosition)
	blackPos := []rune(blackPosition)
	if err := validatePosition(whitePos); err != nil {
		return false, err
	}
	if err := validatePosition(blackPos); err != nil {
		return false, err
	}
	if whitePosition == blackPosition {
		return false, errors.New("queens cannot occupy the same position")
	}
	if blackPos[0] == whitePos[0] || whitePos[1] == blackPos[1] {
		return true, nil
	}
	if math.Abs(float64(whitePos[0]-blackPos[0])) == math.Abs(float64(whitePos[1]-blackPos[1])) {
		return true, nil
	} else {
		return false, nil
	}

}
func validatePosition(position []rune) error {
	if len(position) != 2 {
		return errors.New("invalid position")
	}
	col := position[0]
	row := position[1]
	if col < 'a' || col > 'h' {
		return errors.New("invalid position column")
	}
	if row < '1' || row > '8' {
		return errors.New("invalid position row")
	}
	return nil
}
