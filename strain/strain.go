// Package strain implements a solution for the exercise titled `Strain'.
package strain

// Ints is a slice of int
type Ints []int

// Lists is a slice of Ints
type Lists [][]int

// Strings is a slice of string
type Strings []string

// Keep returns a new collection containing those elements where the predicate is true.
func (xs Ints) Keep(pred func(int) bool) Ints {
	var result Ints
	for _, x := range xs {
		if pred(x) {
			result = append(result, x)
		}
	}
	return result
}

// Discard returns a new collection containing those elements where the predicate is false.
func (xs Ints) Discard(pred func(int) bool) Ints {
	return xs.Keep(func(i int) bool { return !pred(i)})
}

// Keep on Lists (may use Keep on Ints?)
func (xs Lists) Keep(pred func([]int) bool) Lists {
	var result Lists
	for _, x := range xs {
		if pred(x) {
			result = append(result, x)
		}
	}
	return result
}

// Keep on Strings (ditto)
func (xs Strings) Keep(pred func(string) bool) Strings {
	var result Strings
	for _, x := range xs {
		if pred(x) {
			result = append(result, x)
		}
	}
	return result
}
