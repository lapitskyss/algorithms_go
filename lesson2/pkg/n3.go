package pkg

import "fmt"

func N3() {
	var s []int

	from := 3
	to := 20

	num := 0

	s = append(s, from)
	for {
		var news []int

		for _, item := range s {
			a := item + 1
			b := item * 2

			if a == to {
				num++
			} else if a < to {
				news = append(news, a)
			}

			if b == to {
				num++
			} else if b < to {
				news = append(news, b)
			}
		}

		if len(news) == 0 {
			break
		}

		s = news
	}

	fmt.Println(num)
}

func N3V2() {
	from := 3
	to := 20

	num := calc(from, to)

	fmt.Println(num)
}

func calc(from int, to int) int {
	num := 0

	for i := 0; i < 2; i++ {
		n := from
		if i == 0 {
			n++
		} else {
			n = n * 2
		}

		if n < to {
			num += calc(n, to)
		} else {
			if n == to {
				num++
			}
		}
	}

	return num
}
