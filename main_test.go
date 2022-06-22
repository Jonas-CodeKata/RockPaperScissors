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

type testCase struct {
	p1       shape
	p2       shape
	expected result
	msg      string
}

func runTests(t *testing.T, harness ...testCase) {
	for _, test := range harness {
		res := beat(test.p1, test.p2)
		assert.Equal(t, test.expected, res, test.msg)
	}
}

func Test_beat(t *testing.T) {
	tt := []testCase{
		{
			p1:       rock,
			p2:       scissors,
			expected: win,
			msg:      "rock crushes scissors",
		},
		{
			p1:       rock,
			p2:       paper,
			expected: lose,
			msg:      "paper covers rock",
		},
		{
			p1:       rock,
			p2:       rock,
			expected: draw,
			msg:      "rock vs rock is a draw",
		},
	}

	runTests(t, tt...)
}
