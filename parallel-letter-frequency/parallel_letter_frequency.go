// Package letter implements a solution of the exercise
// titled `Parallel Letter Frequency'.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// worker is unexported function, which returns a FreqMap of individual text
// over a channel.
func worker(s string, c chan FreqMap) {
	c <- Frequency(s)
}

// ConcurrentFrequency counts the frequency as function Frequency over
// an array of texts.
func ConcurrentFrequency(texts []string) FreqMap {
	cs := make(chan FreqMap, len(texts))
	for _, text := range texts {
		go worker(text, cs)
	}
	m := FreqMap{}
	for range texts {
		for key, value := range <-cs {
			m[key] += value
		}
	}
	return m
}
