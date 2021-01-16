package pkg

import (
	"fmt"
	"os"
)

func N11() {
	var s []int
	var sum int

	for {
		var n int
		fmt.Print("Введите число: ")
		fmt.Fscan(os.Stdin, &n)

		if n == 0 {
			break
		}

		s = append(s, n)
	}

	for _, num := range s {
		sum += num
	}

	fmt.Println("Среднее арифметическое: ", sum/len(s))
}
