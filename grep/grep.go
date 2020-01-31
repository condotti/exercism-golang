// Package grep implements a solution of the exercise titled `Grep'.
package grep

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func index(slice []string, element string) int {
	for i, member := range slice {
		if member == element {
			return i
		}
	}
	return -1
}

// Search does text search like unix grep command.
func Search(pattern string, flags, files []string) []string {
	nFlag := index(flags, "-n") >= 0
	lFlag := index(flags, "-l") >= 0
	iFlag := index(flags, "-i") >= 0
	vFlag := index(flags, "-v") >= 0
	xFlag := index(flags, "-x") >= 0
	multipleFiles := len(files) > 1

	converter := func(input string) string {
		if iFlag {
			return strings.ToLower(input)
		}
		return input

	}
	collator := func(s, substr string) bool {
		if xFlag {
			return s == substr
		}
		return strings.Contains(s, substr)

	}
	inverter := func(p bool) bool {
		if vFlag {
			return !p
		}
		return p
	}

	result := []string{}
	for _, fileName := range files {
		if file, err := os.Open(fileName); err == nil {
			defer file.Close()
			reader := bufio.NewReaderSize(file, 1024)
			for lineNum := 1; ; lineNum++ {
				line, _, err := reader.ReadLine()
				if err == io.EOF {
					break
				} else if err != nil {
					panic(err)
				}
				if inverter(collator(converter(string(line)), converter(pattern))) {
					if lFlag {
						if index(result, fileName) == -1 {
							result = append(result, fileName)
						}
						continue
					}
					output := string(line)
					if nFlag {
						output = fmt.Sprintf("%v:%v", lineNum, output)
					}
					if multipleFiles {
						output = fmt.Sprintf("%v:%v", fileName, output)
					}
					result = append(result, output)
				}
			}
		} else {
			panic(err)
		}

	}
	return result
}

/*
BenchmarkSearch-4   	      15	  81182687 ns/op	  116170 B/op	    1088 allocs/op
*/
