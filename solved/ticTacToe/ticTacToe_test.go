package main

import "testing"

func TestBoard(t *testing.T) {
	// Board is a string array of 9 length
	arg := board{}
	exp := 9
	res := len(arg)
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
}

func TestCurrentPlayer(t *testing.T) {
	// board.currentPlayer() returns X on a blank board
	arg := board{}
	exp := "X"
	res := arg.currentPlayer()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	// board.currentPlayer() returns O on a board with one X
	arg[0] = "X"
	exp = "O"
	res = arg.currentPlayer()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
}

func TestMove(t *testing.T) {
	// board.move() makes a move on selected index
	arg := board{}
	exp := board{}
	exp[0] = "X"
	arg.move(0)
	res := arg
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	// board.move() can make moves on a board with one X
	exp[1] = "O"
	arg.move(1)
	res = arg
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	// board.move() doesn't allow duplicate moves
	err := arg.move(1)
	res = arg
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	if err == nil {
		t.Fatal("Expected to get an error from move")
	}
}

// board.isOver() returns false for an empty board
func TestIsOverWithEmptyBoard(t *testing.T) {
	arg := board{}
	exp := false
	res := arg.isOver()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
}

// board.isOver() returns true for a cat's game
func TestIsOverWithFullBoard(t *testing.T) {
	arg := board{
		"X", "O", "X",
		"O", "O", "X",
		"O", "X", "O",
	}
	exp := true
	res := arg.isOver()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
}

// board.isOver() returns true when there's a winner
func TestIsOverWithWinner(t *testing.T) {
	// row
	arg := board{
		"X", "X", "X",
		"O", "", "X",
		"O", "", "O",
	}
	exp := true
	res := arg.isOver()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	// col
	arg = board{
		"O", "", "X",
		"O", "", "X",
		"O", "", "O",
	}
	exp = true
	res = arg.isOver()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	// diagonal
	arg = board{
		"O", "", "X",
		"O", "X", "X",
		"X", "", "O",
	}
	exp = true
	res = arg.isOver()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
}

func TestPlayGame(t *testing.T) {
	b := board{}
	b.move(0)
	b.move(1)
	b.move(3)
	b.move(4)
	b.move(6)
	arg := b
	res := b.isOver()
	exp := true
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	resWin := b.winner()
	expWin := "X"
	if resWin != expWin {
		t.Fatalf("Expected %v, got %v with %v arg", expWin, resWin, arg)
	}
}
