package main

import (
	"fmt"
	"unicode"
)

func KingOf(p rune) rune {
	return unicode.ToUpper(p)
}

const (
	Vacant    = ' '
	Red       = 'r'
	RedKing   = 'R'
	Black     = 'b'
	BlackKing = 'B'
)

type Board [8][8]rune

func NewBoard() *Board {
	b := new(Board)

	/*
	 * State Zero
	 */
	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b[0]); col++ {
			b[row][col] = Vacant
		}
	}

	return b
}

func (b *Board) Dump() {
	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b[0]); col++ {
			fmt.Printf("|%c", b[row][col])
		}
		fmt.Println("|")
	}
	fmt.Println("")
	fmt.Println("Red Pieces:", b.ActivePieces(Red))
	fmt.Println("Black Pieces:", b.ActivePieces(Black))
	fmt.Println("")
}

func (b *Board) ActivePieces(color rune) int {
	activePieces := 0
	searchMap := map[rune]bool{
		color:         true,
		KingOf(color): true,
	}

	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b[0]); col++ {
			if searchMap[b[row][col]] {
				activePieces++
			}
		}
	}

	return activePieces
}

func (b *Board) IsGameOver() bool {
	return b.ActivePieces(Red) == 0 || b.ActivePieces(Black) == 0
}
