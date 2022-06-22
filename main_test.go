package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type result = int

const (
	draw result = iota
	win
	lose
)

type shape = int

const (
	rock shape = iota
	paper
	scissors
)

// beat define if the first proposition (prop1) win or lose against the second proposition (prop2),
// if both propositions are the same, the result will be a draw.
func beat(prop1, prop2 shape) result {
	if prop1 == prop2 {
		return draw
	}

	if prop2 == paper {
		return lose
	}
	return win
}

func TestRockBeatScissors(t *testing.T) {
	prop1 := rock
	prop2 := scissors
	res := beat(prop1, prop2)
	assert.Equal(t, win, res, "Rock should beat scissors.")
}

func TestRockLoseAgainstPaper(t *testing.T) {
	prop1 := rock
	prop2 := paper
	res := beat(prop1, prop2)
	assert.Equal(t, lose, res, "Rock should lose against paper.")
}

func TestRockDrawAgainstItSelf(t *testing.T) {
	prop1 := rock
	prop2 := rock
	res := beat(prop1, prop2)
	assert.Equal(t, draw, res, "Rock vs Rock should be draw.")
}

func TestParametricTests(t *testing.T) {
	tt := []struct {
		p1       shape
		p2       shape
		expected result
		msg      string
	}{
		{
			p1:       rock,
			p2:       scissors,
			expected: win,
			msg:      "rock crushes scissors",
		},
	}

	for _, test := range tt {
		res := beat(test.p1, test.p2)
		assert.Equal(t, test.expected, res, test.msg)
	}
}
