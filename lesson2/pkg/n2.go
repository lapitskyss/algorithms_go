package pkg

import "fmt"

func N2() {
	fmt.Println(powV1(2, 5))
}

func powV1(a int, b int) int {
	result := 1

	for i := 0; i < b; i++ {
		result *= a
	}

	return result
}

func N2V2() {
	fmt.Println(powV2(2, 0))
}

func powV2(a int, b int) int {
	if b == 0 {
		return 1
	}

	if b == 1 {
		return a
	}

	if b == 2 {
		return a * a
	}

	return powV2(a, b-1) * a
}
