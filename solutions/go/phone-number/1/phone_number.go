package phonenumber

import (
	"errors"
	"fmt"
)

func Number(phoneNumber string) (string, error) {
	tempnumber := ""
	for _, r := range phoneNumber {
		if r >= '0' && r <= '9' {
			tempnumber += string(r)
		}
	}
	switch {
	case len(tempnumber) < 10 || len(tempnumber) > 11:
		return "", errors.New("invalid number of digits")
	case len(tempnumber) == 11 && tempnumber[0] != '1':
		return "", errors.New("11 digits must start with 1")
	case len(tempnumber) == 11 && tempnumber[0] == '1':
		tempnumber = tempnumber[1:]
	}
	if (tempnumber[0] < '2') || (tempnumber[3] < '2') {
		return "", errors.New("area code cannot start with zero or one")
	}
	return tempnumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return phoneNumber, err
	}
	return string(phoneNumber[0:3]), nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return phoneNumber, err
	}
	return fmt.Sprintf("(%v) %v-%v", string(phoneNumber[0:3]), string(phoneNumber[3:6]), string(phoneNumber[6:])), nil
}
