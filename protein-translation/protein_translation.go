// Package protein implements a solution of the exercise titled `Protein Translation'
// Test code seems somehow strange for error handling... Anyway the code passes the test.
package protein

import (
	"errors"
	"strings"
)

var ErrInvalidBase, ErrStop error

var proteins map[string]string

func init() {
	proteins = map[string]string{}
	for codons, protein := range map[string]string{
		"AUG":                "Methionine",
		"UUU, UUC":           "Phenylalanine",
		"UUA, UUG":           "Leucine",
		"UCU, UCC, UCA, UCG": "Serine",
		"UAU, UAC":           "Tyrosine",
		"UGU, UGC":           "Cysteine",
		"UGG":                "Tryptophan",
		"UAA, UAG, UGA":      "STOP",
	} {
		for _, codon := range strings.Split(codons, ", ") {
			proteins[codon] = protein
		}
	}

	ErrInvalidBase = errors.New("invalid base")
	ErrStop = errors.New("stop sequence")
}

// FromCodon returns a translation of given codon, or an error
func FromCodon(codon string) (string, error) {
	xlated := proteins[codon]
	if xlated == "STOP" {
		xlated = ""
	}
	return xlated, nil
}

// FromRNA returns a series of polypeptide translated from given codons, or an error.
func FromRNA(codons string) ([]string, error) {
	partition := func(s string, n int) []string {
		ar := []rune(s)
		res := []string{}
		for i := 0; i < len(ar)/3; i++ {
			res = append(res, string(ar[i*n:i*n+n]))
		}
		return res
	}
	polypeptide := []string{}
	for _, codon := range partition(codons, 3) {
		xlated, ok := proteins[codon]
		if !ok || xlated == "STOP" {
			return polypeptide, nil
		}
		polypeptide = append(polypeptide, xlated)
	}
	return polypeptide, nil
}
