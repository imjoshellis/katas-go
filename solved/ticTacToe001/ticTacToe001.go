package main

import (
	"errors"
	"fmt"

	"github.com/inancgumus/screen"
)

func main() {
	b := board{}
	var err error
	for !b.isOver() {
		b.print()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Player %v's turn: ", b.currentPlayer())
		var m int
		fmt.Scanln(&m)
		err = b.move(m - 1)
	}
	b.print()
	w := b.winner()
	if w != "" {
		fmt.Printf("Player %v won!\nExiting...", w)
	} else {
		fmt.Printf("It's a tie!\nExiting...")
	}
}

type board [9]string

func (b *board) print() {
	screen.Clear()
	fmt.Print("\n")
	for i, c := range b {
		if c == "" {
			fmt.Printf(" %v ", i+1)
		} else {
			fmt.Printf(" %v ", c)
		}
		if i%3 < 2 {
			fmt.Print("|")
		} else {
			if i < 6 {
				fmt.Println("\n---+---+---")
			}
		}
	}
	fmt.Print("\n\n")
}

func (b *board) currentPlayer() string {
	i := 0
	for _, p := range b {
		if p != "" {
			i++
		}
	}
	if i%2 == 0 {
		return "X"
	}
	return "O"
}

func (b *board) move(i int) error {
	if i < 0 || i > 8 {
		return errors.New("invalid move")
	}
	if b[i] != "" {
		return errors.New("that spot is taken")
	}
	b[i] = b.currentPlayer()
	return nil
}

func (b *board) isOver() bool {
	if b.winner() != "" {
		return true
	}

	for _, c := range b {
		if c == "" {
			return false
		}
	}

	return true
}

func (b *board) winner() string {
	// check rows
	for i := 0; i < 3; i++ {
		if b[i*3] == b[i*3+1] && b[i*3] == b[i*3+2] {
			return b[i*3]
		}
	}

	// check cols
	for i := 0; i < 3; i++ {
		if b[i] == b[i+3] && b[i] == b[i+6] {
			return b[i]
		}
	}

	// check diagonals
	if b[0] == b[4] && b[0] == b[8] {
		return b[0]
	}
	if b[2] == b[4] && b[2] == b[6] {
		return b[2]
	}

	return ""
}
