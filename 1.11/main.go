package main

import "fmt"

func intersections(n1 []int, n2 []int) []int {
	var res []int
	for i := range n1 {
		for j := range n2 {
			if n1[i] == n2[j] {
				res = append(res, n1[i])
			}
		}
	}
	return res
}

func main() {
	n1 := []int{1, 2, 3}
	n2 := []int{2, 3, 4}

	res := intersections(n1, n2)

	fmt.Println(res)
}
