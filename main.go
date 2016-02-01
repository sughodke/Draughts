package main

import (
	"fmt"

	//"board.go"
)

func main() {
	fmt.Printf("%s v%d.%d\n", "Welcome to Go Checkers!", 0, 1)

	b := NewBoard()

	b.ActivePieces(Red)
	b.Dump()

	/*
		rp := NewAutoPlayer(Red)
		bp := NewAutoPlayer(Black)
		for b.IsGameOver() {
			rp.MakeMove(b)
			b.Dump()
			bp.MakeMove(b)
			b.Dump()
		}
	*/
}
