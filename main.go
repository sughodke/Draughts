package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s v%d.%d\n", "Welcome to Go Checkers!", 0, 1)

	b := NewBoard()

	rp := NewPlayer(Red, b)
	bp := NewPlayer(Black, b)

	b.Dump()

	rp.MakeMove(b)
	bp.MakeMove(b)

	/*
		for b.IsGameOver() {
			b.Dump()
			b.Dump()
		}
	*/
}
