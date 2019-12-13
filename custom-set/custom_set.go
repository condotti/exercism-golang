// Package stringset implemnts a solution of the exercise titled `Custom Set'.
package stringset

import "strings"

// Set represnts a custom set with string elements.
type Set map[string]int

// New is a ctor of Set.
func New() Set {
	return Set{}
}

// NewFromSlice is another ctor of Set, initializing with slice of strings.
func NewFromSlice(slice []string) Set {
	s := New()
	for _, elm := range slice {
		s.Add(elm)
	}
	return s
}

// String is the stringer of Set.
func (s Set) String() string {
	elements := []string{}
	for elm, _ := range s {
		elements = append(elements, `"`+elm+`"`)
	}
	return "{" + strings.Join(elements, ", ") + "}"
}

// IsEmpty determines the Set is empty or not.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has determines the Set includes the element or not.
func (s Set) Has(element string) bool {
	_, found := s[element]
	return found
}

// Subset determines Set s1 is a subset of Set s2.
func Subset(s1, s2 Set) bool {
	for elm, _ := range s1 {
		if !s2.Has(elm) {
			return false
		}
	}
	return true
}

// Disjoint determines two Sets are disjoint or not.
func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

// Equal determines two Sets contain same elements or nor.
func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

// Add adds an element to the Set.
func (s Set) Add(element string) {
	s[element] = 0
}

// Intersection returns the intersection of two Sets.
func Intersection(s1, s2 Set) Set {
	s := New()
	for elm, _ := range s1 {
		if s2.Has(elm) {
			s.Add(elm)
		}
	}
	return s
}

// Difference teturns the difference of two Sets.
func Difference(s1, s2 Set) Set {
	s := New()
	for elm, _ := range s1 {
		if !s2.Has(elm) {
			s.Add(elm)
		}
	}
	return s
}

// Union returns the union of two Sets.
func Union(s1, s2 Set) Set {
	s := New()
	for elm, _ := range s1 {
		s.Add(elm)
	}
	for elm, _ := range s2 {
		s.Add(elm)
	}
	return s
}

/*
BenchmarkNewFromSlice1e1-4   	 3100749	       330 ns/op	     256 B/op	       2 allocs/op
BenchmarkNewFromSlice1e2-4   	  121489	      9826 ns/op	    8020 B/op	       9 allocs/op
BenchmarkNewFromSlice1e3-4   	   14306	     82124 ns/op	   61523 B/op	      23 allocs/op
BenchmarkNewFromSlice1e4-4   	    1414	    878131 ns/op	  508066 B/op	     243 allocs/op
*/
