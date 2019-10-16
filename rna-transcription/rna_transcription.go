package strand

func ToRNA(dna string) string {
	pairs := map[string]string { "G":"C", "C":"G", "T":"A", "A":"U" }
	rna := ""
	for _, n := range dna {
		rna += pairs[string(n)]
	}
	return rna
}
