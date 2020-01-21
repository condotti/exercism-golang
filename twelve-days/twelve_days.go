// Package twelve implements a solution of the exercise titled `Twelve Days'.
package twelve

import "strings"

var thing []string = []string{
	"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens",
	"four Calling Birds", "five Gold Rings", "six Geese-a-Laying",
	"seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing",
	"ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming",
}

var nth []string = []string{
	"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}

// Verse returns n'th verse of the song
func Verse(n int) string {
	things := thing[0]
	for i := 1; i < n; i++ {
		if i == 1 {
			things = thing[i] + ", and " + things
		} else {
			things = thing[i] + ", " + things
		}
	}
	return "On the " + nth[n-1] + " day of Christmas my true love gave to me: " + things + "."
}

// Song returns the while lyrics of the song.
func Song() string {
	verse := []string{}
	for i := range thing {
		verse = append(verse, Verse(i+1))
	}
	return strings.Join(verse, "\n")
}

/*
BenchmarkVerse-4   	  164887	      6932 ns/op	   10496 B/op	      78 allocs/op
BenchmarkSong-4    	  144968	     10430 ns/op	   13680 B/op	      84 allocs/op
*/
