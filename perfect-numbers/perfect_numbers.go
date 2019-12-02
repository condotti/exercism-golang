// Package perfect implements a solution of the exercise titled `Perfect Numbers'.
package perfect

import "errors"

// Classification represents classification of natural numbers based on
// Nicomachuss' classification scheme.
type Classification int

// and it's values are defined as follows:
const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

// ErrOnlyPositive will be signaled when input is not positive.
var ErrOnlyPositive error

func init() {
	ErrOnlyPositive = errors.New("only positive input allowed")
}

// Classify returns the classification of input, or error.
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	var aliquotSum, i int64
	for i = 1; i < n; i++ {
		if n%i == 0 {
			aliquotSum += i
		}
	}
	switch {
	case aliquotSum == n:
		return ClassificationPerfect, nil
	case aliquotSum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
