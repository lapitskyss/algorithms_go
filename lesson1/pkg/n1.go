package pkg

import (
	"fmt"
	"os"
)

func N1() {
	var m, h float64

	fmt.Print("Введите вес человека: ")
	fmt.Fscan(os.Stdin, &m)

	fmt.Print("Введите рост человека: ")
	fmt.Fscan(os.Stdin, &h)

	l := m / (h * h)
	fmt.Printf("Индекс массы тела: %g", l)
}
