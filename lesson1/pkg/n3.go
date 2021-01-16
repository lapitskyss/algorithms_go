package pkg

import (
	"fmt"
	"os"
)

func N3() {
	var a, b, c int

	fmt.Print("Введите первое число (a): ")
	fmt.Fscan(os.Stdin, &a)

	fmt.Print("Введите второе число (b): ")
	fmt.Fscan(os.Stdin, &b)

	// первый вариан
	c = a
	a = b
	b = c

	fmt.Printf("a: %d, b: %d", a, b)

	//второй вариан
	a, b = b, a

	fmt.Printf("a: %d, b: %d", a, b)
}
