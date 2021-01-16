package pkg

import (
	"fmt"
	"os"
)

func N9() {
	var n, k int

	fmt.Print("N: ")
	fmt.Fscan(os.Stdin, &n)

	fmt.Print("K: ")
	fmt.Fscan(os.Stdin, &k)

	c := 0
	nn := n

	for n > k {
		c = c + 1
		n = n - k
	}

	fmt.Printf("Частное от деления %d на %d = %d \n", nn, k, c)
	fmt.Printf("Остаток этого деления = %d", n)
}
