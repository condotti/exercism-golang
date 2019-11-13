// Package sublist implements a solution of the exercise titled `sublist'.
package sublist

type Relation string

const (
	equal     Relation = Relation("equal")
	sublist            = Relation("sublist")
	superlist          = Relation("superlist")
	unequal            = Relation("unequal")
)

func partition(n int, l []int) [][]int {
	r := [][]int{}
	for i := 0; i < len(l)-n+1; i++ {
		r = append(r, l[i:i+n])
	}
	return r
}

func deepEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isSublist(a, b []int) bool {
	for _, l := range partition(len(a), b) {
		if deepEqual(a, l) {
			return true
		}
	}
	return false
}

func Sublist(a, b []int) Relation {
	aIsSublistOfB := isSublist(a, b)
	bIsSublistOfA := isSublist(b, a)
	switch {
	case aIsSublistOfB && bIsSublistOfA:
		return equal
	case aIsSublistOfB:
		return sublist
	case bIsSublistOfA:
		return superlist
	default:
		return unequal
	}
}
