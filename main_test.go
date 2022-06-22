package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	draw = iota
	win
	lose
)

func beat(prop1 string, prop2 string) int {
	if prop1 == prop2 {
		return draw
	}

	if prop2 == "paper" {
		return lose
	}
	return win
}

func TestRockBeatScissors(t *testing.T) {
	prop1 := "rock"
	prop2 := "scissors"
	res := beat(prop1, prop2)
	assert.Equal(t, win, res, "Rock should beat scissors.")
}

func TestRockLoseAgainstPaper(t *testing.T) {
	prop1 := "rock"
	prop2 := "paper"
	res := beat(prop1, prop2)
	assert.Equal(t, lose, res, "Rock should lose against paper.")
}

func TestRockDrawAgainstItSelf(t *testing.T) {
	prop1 := "rock"
	prop2 := "rock"
	res := beat(prop1, prop2)
	assert.Equal(t, draw, res, "Rock vs Rock should be draw.")
}
