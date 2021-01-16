package pkg

import (
	"fmt"
	"os"
)

func N2() {
	var n1, n2, n3, n4 int

	fmt.Print("Введите первое число: ")
	fmt.Fscan(os.Stdin, &n1)

	fmt.Print("Введите второе число: ")
	fmt.Fscan(os.Stdin, &n2)

	fmt.Print("Введите третье число: ")
	fmt.Fscan(os.Stdin, &n3)

	fmt.Print("Введите четвертое число: ")
	fmt.Fscan(os.Stdin, &n4)

	max := n1

	if n2 > n1 {
		max = n2
	}

	if n3 > n2 {
		max = n3
	}

	if n4 > n3 {
		max = n4
	}

	fmt.Printf("Максимальное число: %d", max)
}
