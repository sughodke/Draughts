package main

import (
	"fmt"
	"unicode"
)

type Piece rune

func KingOf(p Piece) Piece {
	return unicode.ToUpper(p)
}

const (
	Vacant    = ' '
	Red       = 'r'
	RedKing   = 'R'
	Black     = 'b'
	BlackKing = 'B'
)

type Board [8][8]Piece

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

	for row := 0; row < 3; row++ {
		for i := row % 2; i < 8; i += 2 {
			b[row][(i+1)%8] = Red
			b[7-row][i] = Black
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
}

func (b *Board) ActivePieces(color Piece) int {
	activePieces := 0
	searchMap := map[Piece]bool{
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
