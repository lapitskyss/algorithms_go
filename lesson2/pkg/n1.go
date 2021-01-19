package pkg

import "fmt"

func N1() {
	number := 12

	result := toBin(number)

	fmt.Println(result)
}

func toBin(number int) string {
	var result string

	if number%2 == 0 {
		result = "0" + result
	} else {
		result = "1" + result
	}

	n := number / 2
	if n > 0 {
		result = toBin(n) + result
	}

	return result
}
