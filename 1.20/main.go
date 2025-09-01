package main

import "fmt"

func reverseRunes(r []rune, i int, j int) {
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
}

func reverseWord(s string) string {
	r := []rune(s)

	reverseRunes(r, 0, len(r)-1)

	start := 0

	for i := 0; i <= len(r); i++ {
		if i == len(r) || r[i] == ' ' {
			reverseRunes(r, start, i-1)
			start = i + 1
		}
	}

	return string(r)
}

func main() {
	str := "snow dog sun"
	res := reverseWord(str)

	fmt.Println(str)
	fmt.Println(res)
}
