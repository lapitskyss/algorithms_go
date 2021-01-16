package pkg

import (
	"fmt"
	"os"
)

func N10() {
	var n int

	fmt.Print("Введите число N: ")
	fmt.Fscan(os.Stdin, &n)

	if isHasOddNumber(n) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func isHasOddNumber(n int) bool {
	for {
		num := n % 10
		if num%2 != 0 {
			return true
		}

		n = n / 10

		if n == 0 {
			return false
		}
	}
}
