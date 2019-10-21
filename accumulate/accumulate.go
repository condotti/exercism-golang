package accumulate

// Accumulate takes a COLlection and a CONVerter, and returns a new
// collection containig the result of applying the operation  on
// each element of the input collection.
func Accumulate(col []string, conv func(string) string) []string {
	res := []string{}
	for _, s := range col {
		res = append(res, conv(s))
	}
	return res
}
