package lc

import "strconv"

func fizzBuzzNaive(n int) []string {
	var answer []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if i%3 == 0 {
			answer = append(answer, "Fizz")
		} else if i%5 == 0 {
			answer = append(answer, "Buzz")
		} else {
			answer = append(answer, strconv.Itoa(i))
		}

	}
	return answer
}

func fizzBuzz(n int) []string {
	var answer []string
	for i := 1; i <= n; i++ {
		var str string
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str = strconv.Itoa(i)
		}
		answer = append(answer, str)
	}
	return answer
}

func fizzBuzzHashMap(n int) []string {
	var answer []string
	m := map[int]string{3: "Fizz", 5: "Buzz"}
	keys := []int{3, 5}

	for i := 1; i <= n; i++ {
		var str string

		for _, v := range keys {
			if i%v == 0 {
				str += m[v]
			}
		}

		if str == "" {
			str = strconv.Itoa(i)
		}

		answer = append(answer, str)
	}

	return answer
}
