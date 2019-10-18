package hamming

import "errors"

// Distance returns the hamming distance between a and b, or signals an error
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("length of strands")
	}
	distance := 0
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}
	return distance, nil
}
