package main

import (
	"fmt"
	"tic-tac-toe/board"
)

func main() {
	data := make([]int, 10)
	b := board.Board{data}
	fmt.Println(b.Tokens)
}
