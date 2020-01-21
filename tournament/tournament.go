// Package tournament implemets a solution of exercise titled `Tournament'.
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
		var team [2]*TeamResult
		var ok bool
		for i := range team {
			team[i], ok = tally[cols[i]]
			if !ok {
				team[i] = &TeamResult{team: cols[i]}
				tally[cols[i]] = team[i]
			}
		}
		switch cols[2] {
		case "win":
			team[0].win++
			team[1].lose++
		case "loss":
			team[0].lose++
			team[1].win++
		case "draw":
			team[0].draw++
			team[1].draw++
		default:
			return errors.New("invalid input")
		}
	}
	// convert map to slice
	recs := make([]*TeamResult, 0, len(tally))
	for _, v := range tally {
		v.points += v.win*3 + v.draw
		recs = append(recs, v)
	}
	// sort
	sort.Slice(recs, func(i, j int) bool {
		if recs[i].points == recs[j].points {
			return recs[i].team < recs[j].team
		}
		return recs[j].points < recs[i].points
	})
	// print report
	fmt.Fprintf(b, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	for _, rec := range recs {
		fmt.Fprintf(b, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			rec.team, rec.win+rec.draw+rec.lose, rec.win, rec.draw, rec.lose, rec.points)
	}
	return nil
}

/*
BenchmarkTally-4   	   28579	     39190 ns/op	   43926 B/op	     263 allocs/op
*/
