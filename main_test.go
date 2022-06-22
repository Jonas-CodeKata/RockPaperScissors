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

func beat(prop1 shape, prop2 shape) result {
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
