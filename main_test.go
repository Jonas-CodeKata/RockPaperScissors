package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DrawWhenBothAreTheSame(t *testing.T) {
	runTests(t,
		testCase{
			p1:       rock,
			p2:       rock,
			expected: draw,
			msg:      "rock vs rock give a draw",
		},
		testCase{
			p1:       paper,
			p2:       paper,
			expected: draw,
			msg:      "paper vs paper give a draw",
		},
		testCase{
			p1:       scissors,
			p2:       scissors,
			expected: draw,
			msg:      "scissors vs scissors give a draw",
		},
	)
}

func Test_RockCrushesScissors(t *testing.T) {
	runTests(t,
		testCase{
			p1:       rock,
			p2:       scissors,
			expected: win,
			msg:      "rock crushes scissors",
		},
		testCase{
			p1:       scissors,
			p2:       rock,
			expected: lose,
			msg:      "scissors lose against rock",
		},
	)
}

func Test_PaperCoverRock(t *testing.T) {
	runTests(t,
		testCase{
			p1:       paper,
			p2:       rock,
			expected: win,
			msg:      "paper cover rock",
		},
		testCase{
			p1:       rock,
			p2:       paper,
			expected: lose,
			msg:      "rock lose against paper",
		},
	)
}

func Test_ScissorsCutsPaper(t *testing.T) {
	runTests(t,
		testCase{
			p1:       scissors,
			p2:       paper,
			expected: win,
			msg:      "scissors cut paper",
		},
		testCase{
			p1:       paper,
			p2:       scissors,
			expected: lose,
			msg:      "paper lose against scissors",
		},
	)
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
