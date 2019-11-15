// Package house implements a solution of the exercise titled `House'.
package house

import (
	"fmt"
	"strings"
)

var (
	body        string     = "This is the %s that %s."
	relConn     string     = "%s\nthat %s the %s"
	subjAndVerb [][]string = [][]string{
		{"house"				, "Jack built"},
		{"malt"					, "lay in"},
		{"rat"					, "ate"},
		{"cat"					, "killed"},
		{"dog"					, "worried"},
		{"cow with the crumpled horn"		, "tossed"},
		{"maiden all forlorn"			, "milked"},
		{"man all tattered and torn"		, "kissed"},
		{"priest all shaven and shorn"		, "married"},
		{"rooster that crowed in the morn"	, "woke"},
		{"farmer sowing his corn"		, "kept"},
		{"horse and the hound and the horn"	, "belonged to"},
	}
)

// Verse returns n'th verse of the rhyme.
func Verse(n int) string {
	objSect := subjAndVerb[0][0]
	for i := 0; i < n-1; i++ {
		objSect = fmt.Sprintf(relConn,
			subjAndVerb[i+1][0],
			subjAndVerb[i+1][1],
			objSect)
	}
	return fmt.Sprintf(body, objSect, subjAndVerb[0][1])
}

// Song returns the whole rhyme.
func Song() string {
	rhyme := []string{}
	for i := 0; i < len(subjAndVerb); i++ {
		rhyme = append(rhyme, Verse(i+1))
	}
	return strings.Join(rhyme, "\n\n")
}
