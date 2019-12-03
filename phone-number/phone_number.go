// Package phonenumber implements a solution of the exercise titled `Phone Number'
package phonenumber

import (
	"errors"
	"fmt"
)

// Number returns the number without punctuations or an error if invalid.
func Number(given string) (string, error) {
	err := errors.New("invalid input")
	digits := []rune{}
	for _, digit := range given {
		if digit >= '0' && digit <= '9' {
			digits = append(digits, digit)
		}
	}
	if digits[0] == '1' {
		if len(digits) == 11 {
			digits = digits[1:]
		} else {
			return "", err
		}
	}
	if len(digits) != 10 {
		return "", err
	}
	if digits[0] == '0' || digits[0] == '1' || digits[3] == '0' || digits[3] == '1' {
		return "", err
	}
	return string(digits), nil
}

// AreaCode returns the area code part of given number, or an error if invalid.
func AreaCode(given string) (string, error) {
	number, err := Number(given)
	if err != nil {
		return "", err
	}
	return number[0:3], nil
}

// Format returns formatted phone number, or an error if invalid.
func Format(given string) (string, error) {
	number, err := Number(given)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:]), nil
}
