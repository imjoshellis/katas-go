package main

import "testing"

func TestMove(t *testing.T) {
	b := newBoard()
	arg := 1
	exp := Board("100000000")
	res, err := b.Move(arg)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	arg = 3
	exp = Board("102000000")
	res, err = res.Move(arg)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	arg = 3
	exp = Board("102000000")
	res, err = res.Move(arg)
	if err == nil {
		t.Fatalf("Expected error but got none")
	}
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
}

func TestCurrentPlayer(t *testing.T) {
	arg := newBoard()
	exp := "1"
	res := arg.CurrentPlayer()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
}

func TestPrintBoard(t *testing.T) {
	arg := newBoard()
	exp := " 1 | 2 | 3 \n---+---+---\n 4 | 5 | 6 \n---+---+---\n 7 | 8 | 9 "
	res := arg.printBoard()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}

	arg = Board("010200000")
	exp = " 1 | X | 3 \n---+---+---\n O | 5 | 6 \n---+---+---\n 7 | 8 | 9 "
	res = arg.printBoard()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
}

func TestGameOver(t *testing.T) {
	arg := Board(`121122211`)
	exp := "0"
	res := arg.isOver()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	arg = newBoard()
	exp = ""
	res = arg.isOver()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
	arg = Board("111000000")
	exp = "1"
	res = arg.isOver()
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
}
