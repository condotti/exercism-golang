// Package bowling implements a solution of the exercise titled `Bowling'.
package bowling

import (
	"errors"
)

// Game represents a game of bowling.
type Game struct {
	frame, currentRoll int
	secondThrow        bool
	rolls              [21]int
}

// NewGame is a ctor of Game
func NewGame() *Game {
	return &Game{}
}

// Roll keeps the number of pins knocked down.
func (g *Game) Roll(pins int) error {
	if g.frame > 9 && !g.isStrike(g.currentRoll-2) && !g.isSpare(g.currentRoll-2) {
		return errors.New("Cannot roll after game is over")
	}
	if pins < 0 {
		return errors.New("Negative roll is invalid")
	}
	if pins > 10 {
		return errors.New("Pin count exceeds pins on the lane")
	}
	if g.secondThrow {
		if g.rolls[g.currentRoll-1]+pins > 10 {
			return errors.New("Pin count exceeds pins on the lane")
		}
		g.secondThrow = false
		g.frame++
	} else if pins == 10 {
		g.frame++
	} else {
		g.secondThrow = true
	}
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
	return nil
}

// Score returns the total score of the game.
func (g *Game) Score() (int, error) {
	score := 0
	frameIndex := 0
	strikes := 0
	rollsOf10Frame := 2
	for frame := 0; frame < 10; frame++ {
		if g.isStrike(frameIndex) {
			score += 10 + g.strikeBonus(frameIndex)
			frameIndex++
			if frame < 9 {
				strikes++
			}
			if frame == 9 {
				rollsOf10Frame = 3
			}
		} else if g.isSpare(frameIndex) {
			score += 10 + g.spareBonus(frameIndex)
			frameIndex += 2
			if frame == 9 {
				rollsOf10Frame = 3
			}
		} else {
			score += g.sumOfBallsInFrame(frameIndex)
			frameIndex += 2
		}
	}
	if g.frame < 10 || g.currentRoll < 18-strikes+rollsOf10Frame {
		return 0, errors.New("Score cannot be taken until the end of the game")
	}
	return score, nil
}

func (g *Game) isStrike(frameIndex int) bool {
	return g.rolls[frameIndex] == 10
}

func (g *Game) isSpare(frameIndex int) bool {
	return g.rolls[frameIndex]+g.rolls[frameIndex+1] == 10
}

func (g *Game) sumOfBallsInFrame(frameIndex int) int {
	return g.rolls[frameIndex] + g.rolls[frameIndex+1]
}

func (g *Game) spareBonus(frameIndex int) int {
	return g.rolls[frameIndex+2]
}

func (g *Game) strikeBonus(frameIndex int) int {
	return g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
}
