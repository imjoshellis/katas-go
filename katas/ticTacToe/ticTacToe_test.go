package main

import "testing"

func TestMain(t *testing.T) {
	arg := "argument"
	exp := "expected"
	res := "call function with arg"
	if res != exp {
		t.Fatalf("Expected %v, got %v with %v arg", exp, res, arg)
	}
}
