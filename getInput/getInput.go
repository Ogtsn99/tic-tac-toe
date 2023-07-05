package getInput

import (
	"fmt"
)

func GetPosition() (int, int) {
	fmt.Println("入力してください")
	// コマンドライン引数を取得します
	var x, y int
	fmt.Scan(&x, &y)

	return x, y
}

func Validate(value1 int, value2 int) bool {
	if (value1 == 0 || value1 == 1 || value1 == 2) && (value2 == 0 || value2 == 1 || value2 == 2) {
		return true
	} else {
		return false
	}
}
