package main

import (
	"errors"
	"fmt"
)

type Location struct {
	row, col int
}

func (lhs *Location) Add(rhs Location) Location {
	return Location{row: lhs.row + rhs.row, col: lhs.col + rhs.col}
}

type Player struct {
	color  rune
	pieces map[rune]Location
}

func NewPlayer(col rune, b *Board) *Player {
	p := Player{color: col}
	p.pieces = make(map[rune]Location)

	for row := 0; row < 3; row++ {
		for i := row % 2; i < 8; i += 2 {
			var x, y int

			switch col {
			case Red:
				x = row
				y = (i + 1) % 8
			case Black:
				x = 7 - row
				y = i
			}

			b[x][y] = col
			p.pieces['0'+rune(len(p.pieces))] = Location{row: x, col: y}
		}
	}

	return &p
}

func (p *Player) moveLeft(piece rune, b *Board) {
	/*
	 *  var factor int
	 *
	 *  switch p.color {
	 *  case Red, RedKing:
	 *    factor = 1
	 *  case Black, BlackKing:
	 *    factor = -1
	 *  }
	 */

	location := p.pieces['b'-piece]
	jump := Location{-1, -1}
	jumpLocation := location.Add(jump)
	fmt.Println("", jumpLocation)
	// check if there is an enemy piece
	//b.WhatsOn(jumpLocation) is in Enemy Map,
	// Add another Jump
	// jumpLocation := location.Add(jump)
	// check if the move is still on the board
}

func (p *Player) MakeMove(b *Board) {
	var input string

	fmt.Scanln(&input)

	piece := input[0]
	if piece < '0' || piece > 'b' {
		errors.New("Selected an unknown piece")
	}

	for i := 1; i < len(input); i++ {
		move := input[i]

		switch move {
		case 'l':
			p.moveLeft(rune(piece), b)
		case 'r':
			// p.moveRight(b)
		}
	}
}
