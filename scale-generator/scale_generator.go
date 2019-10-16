package scale

import (
	"strings"
)

func IndexOf(elm string, arr []string) int {
	for ind, e := range arr {
		if e == elm {
			return ind
		}
	}
	return -1		// not found
}

func Scale(tonic, interval string) []string {
	chroma := [][]string {
		{ "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#" },
		{ "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab" },
	}
	tonics := [][]string {
		{ "C", "G", "A", "E", "B", "F#", "e", "b", "a", "f#", "c#", "g#", "d#" },
		{ "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb" },
	}
	intervals := []string { "m", "M", "A" }
	var scale []string

	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}
	for i := 0; i < 2; i++ {
		if IndexOf(tonic, tonics[i]) >= 0 {
			scale = append(scale, strings.Title(tonic))
			idx := IndexOf(strings.Title(tonic), chroma[i])
			for j := 0; j < len(interval) - 1; j++ {
				idx += IndexOf(interval[j:j + 1], intervals) + 1
				scale = append(scale, strings.Title(chroma[i][idx % len(chroma[i])]))
			}
		}
	}
	return scale		// nil
}
