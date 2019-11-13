// Package flatten implements a solution for the exercise titled `Flatten Array'.
package flatten

// Flatten takes a nested list and returns a single flattened list with all values except nil/null.
func Flatten(xs interface{}) []interface{} {
	r := []interface{}{}
	switch x := xs.(type) {
	case nil: // nop
	case []interface{}:
		for _, y := range x {
			r = append(r, Flatten(y)...)
		}
	default:
		r = append(r, x)
	}
	return r
}
