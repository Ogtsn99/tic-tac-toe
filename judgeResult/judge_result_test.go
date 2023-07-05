package judgeResult

import (
	"testing"
	"tic-tac-toe/board"
)

func Test_JudgeResult(t *testing.T) {
	expected := 0
	data := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	b := board.Board{Tokens: data}

	if expected != JudgeResult(b) {
		t.Errorf("Test fail expected: 0")
	}
}
