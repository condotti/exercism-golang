package letter

import "sync"
import "fmt"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	fmt.Println(m)
	return m
}

// safeFreqMap is unexported struct, contains FreqMap and a mutex
type safeFreqMap struct {
	m FreqMap;
	mux sync.Mutex
	wg sync.WaitGroup
}

// frequency is unexported method, counts like Frequency with mutex operation.
func (m *safeFreqMap)frequency(s string) {
	m.wg.Add(1)
	for _, r := range s {
		m.mux.Lock()
		m.m[r]++
		m.mux.Unlock()
	}
	m.wg.Done()
	fmt.Println(m.m)
}

// ConcurrentFrequency counts the frequency as function Frequency over
// array of texts.
func ConcurrentFrequency(xs []string) FreqMap {
	m := safeFreqMap{m: FreqMap{}}
	for _, text := range xs {
		go m.frequency(text)
	}
	m.wg.Wait()
	fmt.Println(m.m)
	return m.m
}
