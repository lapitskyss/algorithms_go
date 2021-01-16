package pkg

import (
	"fmt"
	"os"
)

func N7() {
	var x1, y1, x2, y2 int

	fmt.Print("x1: ")
	fmt.Fscan(os.Stdin, &x1)

	fmt.Print("y1: ")
	fmt.Fscan(os.Stdin, &y1)

	fmt.Print("x2: ")
	fmt.Fscan(os.Stdin, &x2)

	fmt.Print("y2: ")
	fmt.Fscan(os.Stdin, &y2)

	if getColor(x1, y1) == getColor(x2, y2) {
		fmt.Print("Поля относятся к одному цвету")
	} else {
		fmt.Print("Поля НЕ относятся к одному цвету")
	}
}

func getColor(x int, y int) string {
	if (x%2 == 0 && y%2 == 0) || (x%2 != 0 && y%2 != 0) {
		return "black"
	}

	return "white"
}
