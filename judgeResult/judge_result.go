package judgeResult

import "tic-tac-toe/board"

// Boardをみて誰かが勝っているかを判定するプログラム
// 決着がついていないなら0, player1()が勝っているなら1, player2が勝っているなら2
func JudgeResult(b board.Board) int {
	if b.Tokens[0] == 2 {
		if b.Tokens[1] == 2 {
			if b.Tokens[2] == 2 {
				return 2
			}
		}
	}

	if b.Tokens[0] == 1 {
		if b.Tokens[1] == 1 {
			if b.Tokens[2] == 1 {
				return 1
			}
		}
	}

	if b.Tokens[3] == 2 {
		if b.Tokens[4] == 2 {
			if b.Tokens[5] == 2 {
				return 2
			}
		}
	}

	if b.Tokens[3] == 1 {
		if b.Tokens[4] == 1 {
			if b.Tokens[5] == 1 {
				return 1
			}
		}
	}

	if b.Tokens[6] == 2 {
		if b.Tokens[7] == 2 {
			if b.Tokens[8] == 2 {
				return 2
			}
		}
	}

	if b.Tokens[6] == 1 {
		if b.Tokens[7] == 1 {
			if b.Tokens[8] == 1 {
				return 1
			}
		}
	}

	if b.Tokens[0] == 2 {
		if b.Tokens[3] == 2 {
			if b.Tokens[6] == 2 {
				return 2
			}
		}
	}

	if b.Tokens[0] == 1 {
		if b.Tokens[3] == 1 {
			if b.Tokens[6] == 1 {
				return 1
			}
		}
	}

	if b.Tokens[1] == 2 {
		if b.Tokens[4] == 2 {
			if b.Tokens[7] == 2 {
				return 2
			}
		}
	}

	if b.Tokens[1] == 1 {
		if b.Tokens[4] == 1 {
			if b.Tokens[7] == 1 {
				return 1
			}
		}
	}

	if b.Tokens[2] == 2 {
		if b.Tokens[5] == 2 {
			if b.Tokens[8] == 2 {
				return 2
			}
		}
	}

	if b.Tokens[2] == 1 {
		if b.Tokens[5] == 1 {
			if b.Tokens[8] == 1 {
				return 1
			}
		}
	}

	if b.Tokens[0] == 2 {
		if b.Tokens[4] == 2 {
			if b.Tokens[8] == 2 {
				return 2
			}
		}
	}

	if b.Tokens[0] == 1 {
		if b.Tokens[4] == 1 {
			if b.Tokens[8] == 1 {
				return 1
			}
		}
	}

	if b.Tokens[2] == 2 {
		if b.Tokens[4] == 2 {
			if b.Tokens[6] == 2 {
				return 2
			}
		}
	}

	if b.Tokens[2] == 1 {
		if b.Tokens[4] == 1 {
			if b.Tokens[6] == 1 {
				return 1
			}
		}
	}

	return 0

}
