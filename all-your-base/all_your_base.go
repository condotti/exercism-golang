// Package allyourbase implements a solution of the exercise titled `All Your Base'.
package allyourbase

import "errors"

// ConvertToBase converts a number to any other base.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	number := 0
	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		number *= inputBase
		number += digit
	}
	outputDigits := []int{number % outputBase}
	for number /= outputBase; number > 0; number /= outputBase {
		outputDigits = append([]int{number % outputBase}, outputDigits...)
	}
	return outputDigits, nil
}
