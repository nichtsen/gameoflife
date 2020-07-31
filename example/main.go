package main

import(
	"github.com/nichtsen/gameoflife"
)

func main() {
	//board length, width,
	//and block length by pixels
	width, len, unit := 5, 5, 100
	//1 for alive 0 for dead
	b, err := newBlock(
		width,
		len,
		[]byte{
		0,1,0,0,1,
		0,1,1,1,0,
		0,1,1,0,0,
		0,1,0,0,0,
		1,0,0,1,0},
	)
	if err != nil {
		panic(err)
	}
	br := newBoard(&b, unit)
	// transition for 10 times
	br.transition(10)
}


