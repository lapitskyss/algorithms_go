package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := randomArray()

	fmt.Println("Sergey Lapitsky")

	// Пузырьковая сортровка
	fmt.Println("1. Попробовать оптимизировать пузырьковую сортировку. Сравнить количество операций сравнения оптимизированной и не оптимизированной программы. Написать функции сортировки, которые возвращают количество операций.")
	fmt.Println(bubbleSort(arr))

	//  Шейкерная сортировка
	fmt.Println("2. *Реализовать шейкерную сортировку.")
	fmt.Println(cocktailSort(arr))

	// Бинарный алгоритм поиска
	fmt.Println("3. Реализовать бинарный алгоритм поиска в виде функции, которой передается отсортированный массив. Функция возвращает индекс найденного элемента или -1, если элемент не найден.")
	fmt.Println(binSearch(cocktailSort(arr), 8))
}

func randomArray() []int {
	list := rand.Perm(20)
	for i := 0; i < len(list); i++ {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	return list
}

func bubbleSort(arr []int) []int {
	var n = len(arr)

	newArr := make([]int, len(arr))
	copy(newArr, arr)

	for {
		newn := 0
		for i := 1; i < n; i++ {
			if newArr[i-1] > newArr[i] {
				newArr[i-1], newArr[i] = newArr[i], newArr[i-1]
				newn = i
			}
		}
		n = newn
		if n <= 1 {
			break
		}
	}

	return newArr
}

func cocktailSort(arr []int) []int {
	var n = len(arr)

	newArr := make([]int, len(arr))
	copy(newArr, arr)

	for i := 0; i < (n - 1); i++ {
		for j := n - 1; j > i; j-- {
			if newArr[j] < newArr[j-1] {
				newArr[j], newArr[j-1] = newArr[j-1], newArr[j]
			}
		}
	}

	return newArr
}

func binSearch(arr []int, item int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + high
		guess := arr[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
