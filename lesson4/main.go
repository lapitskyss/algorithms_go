package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sergey Lapitsky")

	fmt.Println("2. Решить задачу о нахождении длины максимальной последовательности с помощью матрицы.")

	x := []int{1, 5, 2, 8, 9, 7, 4}
	y := []int{1, 5, 9, 2, 1, 6, 9, 7}
	fmt.Println(largestSubsequence(x, y))
}

func largestSubsequence(x, y []int) int {
	l := make(map[int]map[int]int)

	for xI, xElem := range x {
		if l[xI] == nil {
			l[xI] = map[int]int{}
		}

		for yI, yElem := range y {
			if xElem == yElem {
				l[xI][yI] = l[xI-1][yI-1] + 1
			} else {
				l[xI][yI] = max(l[xI][yI-1], l[xI-1][yI])
			}
		}
	}

	return l[len(x)-1][len(y)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
