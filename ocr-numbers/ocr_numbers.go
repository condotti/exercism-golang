// Package ocr implements a solution of the exercise titled `OCR Numbers'.
package ocr

import (
	"strconv"
	"strings"
)

func partition(s string, l int) []string {
	parted := []string{}
	for ; len(s) >= l; s = s[l:] {
		parted = append(parted, s[:l])
	}
	return parted
}

func chunkOff(s []string) []string {
	chunks := []string{}
	for _, line := range s {
		for i, chunk := range partition(line, 3) {
			if len(chunks) <= i {
				chunks = append(chunks, "")
			}
			chunks[i] += chunk
		}
	}
	return chunks
}

var pattern []string = chunkOff(strings.Split(`
 _     _  _     _  _  _  _  _ 
| |  | _| _||_||_ |_   ||_||_|
|_|  ||_  _|  | _||_|  ||_| _|
                              `, "\n"))

func recognizeLine(s []string) string {
	digits := ""
	for _, chunk := range chunkOff(s) {
		digits += func(s string) string {
			for i, pat := range pattern {
				if s == pat {
					return strconv.Itoa(i)
				}
			}
			return "?"
		}(chunk)
	}
	return digits
}

// RecognizeDigit is a dummy function (test requires)
func recognizeDigit(s string) string {
	return ""
}

// Recognize returns a recognized ocr number as a string.
func Recognize(in string) []string {
	out := []string{}
	lines := strings.Split(in, "\n")[1:]
	for len(lines) > 0 {
		out = append(out, recognizeLine(lines[:4]))
		lines = lines[4:]
	}
	return out
}
