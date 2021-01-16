package pkg

import (
	"fmt"
	"os"
)

func N12() {
	var n1, n2, n3, max int

	fmt.Print("Введите первое число: ")
	fmt.Fscan(os.Stdin, &n1)

	fmt.Print("Введите второе число: ")
	fmt.Fscan(os.Stdin, &n2)

	fmt.Print("Введите третье число: ")
	fmt.Fscan(os.Stdin, &n3)

	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}

	if n3 > max {
		max = n3
	}

	fmt.Printf("Максимальное число: %d", max)
}
