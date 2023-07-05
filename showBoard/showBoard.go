package judgeResult

import (
	"fmt"
)

func ShowBoard(b [9]int) {
	for i, v := range b {
		if v == 0 {
			// empty space. Display number
			fmt.Printf(".")
		} else if v == 1 {
			fmt.Printf("X")
		} else if v == 2 {
			fmt.Printf("O")
		}
		// And now the decorations
		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" | ")
		}
	}
}
