package hamming

import "errors"

// Distance returns the hamming distance between a and b, or signals an error
func Distance(a, b string) (int, error) {
	if (len(a) != len(b)) {
		return 0, errors.New("length of strands")
	}
	distance := 0
	for i := range a {
		if a[i:i+1] != b[i:i+1] {
			distance++
		}
	}
	return distance, nil
}
