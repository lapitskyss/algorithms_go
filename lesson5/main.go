package main

import "fmt"

func main() {

	fmt.Println("Sergey Lapitsky")

	fmt.Println("1. Реализовать перевод из десятичной в двоичную систему счисления с использованием стека.")
	fmt.Println(fromDecimalToBinary(123))

	fmt.Println("3. Написать программу, которая определяет, является ли введенная скобочная последовательность правильной. Примеры правильных скобочных выражений: (), ([])(), {}(), ([{}]), неправильных — )(, ())({), (, ])}), ([(]) для скобок [,(,{.")
	fmt.Println(isParenthesisSequenceCorrect("(2+(2*2)))"))
}

func fromDecimalToBinary(number int) string {
	var stack []string

	for {
		if number%2 == 0 {
			stack = append(stack, "0")
		} else {
			stack = append(stack, "1")
		}

		number = number / 2

		if number <= 0 {
			break
		}
	}

	result := ""

	for len(stack) > 0 {
		n := len(stack) - 1
		result += stack[n]

		stack = stack[:n]
	}

	return result
}

func isParenthesisSequenceCorrect(str string) bool {
	var stack Stack
	success := true

	for _, r := range str {
		brace := string(r)
		if checkLeftBrace(brace) {
			stack.Push(brace)
		} else if checkRightBrace(brace) {
			if stack.IsEmpty() {
				success = false
				break
			}
			pair, _ := stack.Pop()
			if getTheBraceLeftPair(brace) != pair {
				success = false
				break
			}
		}
	}

	return success
}

func checkLeftBrace(brace string) bool {
	return brace == "(" || brace == "{" || brace == "["
}

func checkRightBrace(brace string) bool {
	return brace == ")" || brace == "}" || brace == "]"
}

func getTheBraceLeftPair(brace string) string {
	if brace == ")" {
		return "("
	} else if brace == "}" {
		return "{"
	} else if brace == "]" {
		return "["
	}

	return brace
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
