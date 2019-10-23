// Pakcage letter implements a solution of the exercise
// titled `Parallel Letter Frequency'.
// Unfortunately this solution passes and failes the test case randomely
// (so far it seems randomely, i mean).
package letter

import "sync"

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

// safeFreqMap is unexported struct, contains FreqMap and a mutex
type safeFreqMap struct {
	m FreqMap;
	sync.Mutex
	sync.WaitGroup
}

// frequency is unexported method, counts like Frequency with mutex operation.
func (m *safeFreqMap)frequency(s string) {
	m.Add(1)
	defer m.Done()
	for _, r := range s {
		m.Lock()
		m.m[r]++
		m.Unlock()
	}
}

// ConcurrentFrequency counts the frequency as function Frequency over
// array of texts.
func ConcurrentFrequency(xs []string) FreqMap {
	m := safeFreqMap{m: FreqMap{}}
	for _, text := range xs {
		go m.frequency(text)
	}
	m.Wait()
	return m.m
}
