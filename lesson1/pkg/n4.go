package pkg

import (
	"fmt"
	"math"
	"os"
)

func N4() {
	var a, b, c float64

	fmt.Print("x2 − 2x − 3 = 0 ⇒ a = 1; b = −2; c = −3;")

	fmt.Print("Введите a: ")
	fmt.Fscan(os.Stdin, &a)

	fmt.Print("Введите b: ")
	fmt.Fscan(os.Stdin, &b)

	fmt.Print("Введите c: ")
	fmt.Fscan(os.Stdin, &c)

	d := math.Pow(b, 2) - 4*a*c

	if d < 0 {
		fmt.Print("Корней нет")
	}

	if d == 0 {
		x1 := (-b - math.Sqrt(d)) / 2 * a
		fmt.Printf("У уравнения один корень, %g", x1)
	}

	if d > 0 {
		x1 := (-b - math.Sqrt(d)) / 2 * a
		x2 := (-b + math.Sqrt(d)) / 2 * a

		fmt.Printf("У уравнения два кореня, %g %g", x1, x2)
	}
}
