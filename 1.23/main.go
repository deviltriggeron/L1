package main

import "fmt"

func deleteByIndex(arr []int, index int) []int {
	if index < 0 || index > len(arr) {
		return arr
	}

	copy(arr[index:], arr[index+1:])

	var zero int
	arr[len(arr)-1] = zero

	return arr[:len(arr)-1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	res := deleteByIndex(arr, 2)

	fmt.Println(res)
}
