package main

import (
	"fmt"
	"os"
	"tic-tac-toe/board"
	"tic-tac-toe/getInput"
	"tic-tac-toe/judgeResult"
	judgeResult2 "tic-tac-toe/showBoard"
)

func main() {
	data := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b := board.Board{Tokens: data}
	judgeResult2.ShowBoard(b.Tokens)
	turn := 1
	for judgeResult.JudgeResult(b) == 0 {
		if turn == 1 {
			fmt.Println("player1 入力してください")
		} else {
			fmt.Println("player2 入力してください")
		}

		x, y := getInput.GetPosition()
		judgeResult2.ShowBoard(b.Tokens)
		if turn == 1 {
			b.Tokens[x*3+y] = 1
		} else {
			b.Tokens[x*3+y] = 2
		}

		println("======")
		judgeResult2.ShowBoard(b.Tokens)

		canContinue := false
		for _, v := range b.Tokens {
			if v == 0 {
				canContinue = true
			}
		}

		if !canContinue {
			fmt.Println("引き分けです")
			os.Exit(0)
		}

		turn = 1 - turn
	}

	if judgeResult.JudgeResult(b) == 1 {
		fmt.Println("Player1が勝ちました！！")
	} else {
		fmt.Println("")
	}
}
