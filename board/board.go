package board

type Board struct {
	Tokens [9]int
}

func (b *Board) Put(x int, y int, val int) bool {
	if b.Tokens[x*3+y] != 0 {
		return false
	}
	b.Tokens[x*3+y] = val
	return true
}
