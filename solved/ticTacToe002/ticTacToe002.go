package main

import (
	"errors"
	"fmt"
)

type Board string

func main() {
	b := newBoard()
	var err error
	for b.isOver() == "" {
		fmt.Println("\n")
		fmt.Println(b.printBoard())
		var n int
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Println("Choose a spot by entering its number")
		fmt.Printf("Player %s's turn: ", b.CurrentPlayer())
		fmt.Scanln(&n)
		b, err = b.Move(n)
		if b.isOver() != "" {
			fmt.Println("\n")
			fmt.Println(b.printBoard())
			fmt.Printf("Game Over: Player %s won!\n", b.isOver())
		}
	}
	fmt.Printf("Press enter to exit")
	fmt.Scanln()
}

func (b Board) Move(p int) (Board, error) {
	if p > 9 || p < 1 {
		return b, errors.New("invalid move")
	}
	p--
	if b[p] != '0' {
		return b, errors.New("that spot is taken")
	}
	b = b[:p] + Board(b.CurrentPlayer()) + b[p+1:]
	return b, nil
}

func (b Board) CurrentPlayer() string {
	p1 := 0
	p2 := 0
	for _, c := range b {
		if c == '1' {
			p1++
		}
		if c == '2' {
			p2++
		}
	}
	if p1 == p2 {
		return "1"
	} else {
		return "2"
	}
}

func newBoard() Board {
	return "000000000"
}

func (b Board) printBoard() string {
	s := ""
	for i, c := range b {
		if c == '0' {
			s += fmt.Sprintf(" %d ", i+1)
		}
		if c == '1' {
			s += fmt.Sprintf(" %s ", "X")
		}
		if c == '2' {
			s += fmt.Sprintf(" %s ", "O")
		}
		if i%3 < 2 {
			s += "|"
		}
		if i%3 == 2 && i != 8 {
			s += "\n---+---+---\n"
		}
	}

	return s
}

func (b Board) isOver() string {
	for i := 0; i < 3; i++ {
		if b[0+i*3] == '0' {
			continue
		}
		if b[0+i*3] == b[1+i*3] && b[0+i*3] == b[2+i*3] {
			return string(b[0+i*3])
		}
	}
	for i := 0; i < 3; i++ {
		if b[0+i] == '0' {
			continue
		}
		if b[0+i] == b[3+i] && b[0+i] == b[6+i] {
			return string(b[0+i])
		}
	}
	if b[0] != '0' && b[0] == b[4] && b[0] == b[8] {
		return string(b[0])
	}
	if b[2] != '0' && b[2] == b[4] && b[2] == b[6] {
		return string(b[2])
	}
	for _, c := range b {
		if c == '0' {
			return ""
		}
	}
	return "0"
}
