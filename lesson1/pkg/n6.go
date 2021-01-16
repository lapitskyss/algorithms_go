package pkg

import (
	"fmt"
	"os"
)

func N6() {
	var age int

	fmt.Print("Введите возраст человека (от 1 до 150 лет): ")
	fmt.Fscan(os.Stdin, &age)

	var txt string
	count := age % 100
	if count >= 5 && count <= 20 {
		txt = "лет"
	} else {
		count = count % 10
		if count == 1 {
			txt = "год"
		} else if count >= 2 && count <= 4 {
			txt = "года"
		} else {
			txt = "лет"
		}
	}

	fmt.Printf("%d %s", age, txt)
}
