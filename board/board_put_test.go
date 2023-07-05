package board

import (
	"testing"
)

func Test_BoardCanPut(t *testing.T) {
	data := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	expected := [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	board := Board{data}

	if true != board.Put(0, 0, 1) {
		t.Errorf("Test fail expected: %s", "true")
	}

	if expected != board.Tokens {
		t.Errorf("Test fail expected: 1, 0, 0, 0, 0, 0, 0, 0, 0")
	}
}

func Test_BoardCanNotPut(t *testing.T) {
	data := [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	expected := [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	board := Board{data}

	if false != board.Put(0, 0, 1) {
		t.Errorf("Test fail expected: %s", "true")
	}

	if expected != board.Tokens {
		t.Errorf("Test fail expected: {1, 0, 0, 0, 0, 0, 0, 0, 0}")
	}
}
