package pkg

import (
	"fmt"
	"os"
)

func N8() {
	var a, b int

	fmt.Print("a: ")
	fmt.Fscan(os.Stdin, &a)

	fmt.Print("b: ")
	fmt.Fscan(os.Stdin, &b)

	if b < a {
		fmt.Println("b < a")
	}

	for {
		if a > b {
			break
		}

		fmt.Printf("Квадрат %d: %d \n", a, a*a)
		fmt.Printf("Куб %d: %d \n", a, a*a*a)
		a++
	}
}
