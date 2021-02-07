package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sergey Lapitsky")

	fmt.Println("1. Реализовать простейшую хеш-функцию. На вход функции подается строка, на выходе сумма кодов символов.")
	fmt.Println(getHash("qwecdw23fd3f"))
}

func getHash(str string) int {
	var sum int
	for _, r := range str {
		sum += int(r)
	}

	return sum
}
