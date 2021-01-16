package pkg

import (
	"fmt"
	"os"
)

func N5() {
	var month int

	fmt.Print("Введите номер месяца: ")
	fmt.Fscan(os.Stdin, &month)

	switch {
	case month <= 2 || month == 12:
		fmt.Println("Зима")
	case month >= 3 || month <= 5:
		fmt.Println("Весна")
	case month >= 6 || month <= 8:
		fmt.Println("Лето")
	case month >= 9 || month <= 11:
		fmt.Println("Осень")
	}
}
