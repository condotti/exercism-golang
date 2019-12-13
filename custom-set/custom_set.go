// Package stringset implemnts a solution of the exercise titled `Custom Set'.
package stringset

import "strings"

// Set represnts a custom set with string elements.
type Set []string

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
	// fmt.Println("nfs", s)
	return s
}

// String is the stringer of Set.
func (s Set) String() string {
	slice := []string{}
	for _, elm := range s {
		slice = append(slice, `"`+elm+`"`)
	}
	return "{" + strings.Join(slice, ", ") + "}"
}

// IsEmpty determines the Set is empty or not.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has determines the Set includes the element or not.
func (s Set) Has(element string) bool {
	for _, elm := range s {
		if elm == element {
			return true
		}
	}
	return false
}

// Subset determines Set s1 is a subset of Set s2.
func Subset(s1, s2 Set) bool {
	for _, element := range s1 {
		if !s2.Has(element) {
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
func (s *Set) Add(element string) {
	if !s.Has(element) {
		*s = append(*s, element)
	}
	// fmt.Println(s)
}

// Intersection returns the intersection of two Sets.
func Intersection(s1, s2 Set) Set {
	s := New()
	for _, elm := range s1 {
		if s2.Has(elm) {
			s.Add(elm)
		}
	}
	return s
}

// Difference teturns the difference of two Sets.
func Difference(s1, s2 Set) Set {
	s := New()
	for _, elm := range s1 {
		if !s2.Has(elm) {
			s.Add(elm)
		}
	}
	return s
}

// Union returns the union of two Sets.
func Union(s1, s2 Set) Set {
	return NewFromSlice(append(s1, s2...))
}

/*
BenchmarkNewFromSlice1e1-4   	 2587545	       446 ns/op	     240 B/op	       4 allocs/op
BenchmarkNewFromSlice1e2-4   	   71619	     14984 ns/op	    2032 B/op	       7 allocs/op
BenchmarkNewFromSlice1e3-4   	     877	   1410185 ns/op	   32752 B/op	      11 allocs/op
BenchmarkNewFromSlice1e4-4   	       8	 140751125 ns/op	  473712 B/op	      18 allocs/op
*/
