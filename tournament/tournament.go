// package tournament implemets a solution of exercise titled `Tournament'.
// I referred to the document of the package csv
// (https://golang.org/pkg/encoding/csv/#example_Reader_options).

package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
)

// TeamResult contains the name of team,  w/d/l count and points of each team
type TeamResult struct {
	team                    string
	win, draw, lose, points int
}

// ByTeam is for sort
type ByTeam []TeamResult

func (a ByTeam) Len() int      { return len(a) }
func (a ByTeam) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTeam) Less(i, j int) bool {
	if a[i].points == a[j].points {
		return a[i].team < a[j].team
	}
	return a[j].points < a[i].points
}

// Tally tallies the input containig results of competitions.
func Tally(sr io.Reader, b io.Writer) error {
	// read at once
	r := csv.NewReader(sr)
	r.Comma = ';'
	r.Comment = '#'
	lines, err := r.ReadAll()
	if err != nil {
		return err
	}
	// tally up
	tally := map[string]*TeamResult{}
	for _, cols := range lines {
		if len(cols) != 3 {
			return errors.New("invalid input")
		}
		team0, ok := tally[cols[0]]
		if !ok {
			team0 = &TeamResult{cols[0], 0, 0, 0, 0}
			tally[cols[0]] = team0
		}
		team1, ok := tally[cols[1]]
		if !ok {
			team1 = &TeamResult{cols[1], 0, 0, 0, 0}
			tally[cols[1]] = team1
		}
		switch cols[2] {
		case "win":
			team0.win++
			team1.lose++
		case "loss":
			team0.lose++
			team1.win++
		case "draw":
			team0.draw++
			team1.draw++
		default:
			return errors.New("invalid input")
		}
	}
	recs := ByTeam{}
	for _, v := range tally {
		v.points += v.win*3 + v.draw
		recs = append(recs, *v)
	}
	sort.Sort(ByTeam(recs))
	fmt.Fprintf(b, "Team                           | MP |  W |  D |  L |  P\n")
	for _, rec := range recs {
		fmt.Fprintf(b, "%-30s | %2d | %2d | %2d | %2d | %2d\n", rec.team, rec.win+rec.draw+rec.lose, rec.win, rec.draw, rec.lose, rec.points)
	}
	return nil
}
