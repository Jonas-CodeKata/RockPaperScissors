package main

type result int

const (
	draw result = iota
	win
	lose
)

type shape int

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

	if prop1 == rock && prop2 == scissors ||
		prop1 == scissors && prop2 == paper ||
		prop1 == paper && prop2 == rock {
		return win
	}

	return lose
}
