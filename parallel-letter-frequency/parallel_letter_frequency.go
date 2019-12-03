// Package letter implements a solution of the exercise
// titled `Parallel Letter Frequency'.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	fm := FreqMap{}
	for _, letter := range text {
		fm[letter]++
	}
	return fm
}

// ConcurrentFrequency counts the frequency as function Frequency over
// an array of texts.
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap, len(texts))
	for _, text := range texts {
		go func(text string, c chan FreqMap) {
			c <- Frequency(text)
		}(text, c)
	}
	fm := FreqMap{}
	for range texts {
		for letter, freq := range <-c {
			fm[letter] += freq
		}
	}
	return fm
}
