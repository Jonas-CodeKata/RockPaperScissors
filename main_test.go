package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const draw = "draw"

func beat(prop1 string, prop2 string) bool {
	if prop2 == "paper" {
		return false
	}
	return true
}

func TestRockBeatScissors(t *testing.T) {
	prop1 := "rock"
	prop2 := "scissors"
	res := beat(prop1, prop2)
	assert.True(t, res, "Rock should beat scissors.")
}

func TestRockLoseAgainstPaper(t *testing.T) {
	prop1 := "rock"
	prop2 := "paper"
	res := beat(prop1, prop2)
	assert.False(t, res, "Rock should lose against paper.")
}

func TestRockDrawAgainstItSelf(t *testing.T) {
	prop1 := "rock"
	prop2 := "rock"
	res := beat(prop1, prop2)
	assert.Equal(t, draw, res, "Rock vs Rock should be draw.")
}
