package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	first := arr[0]
	var left, right []int

	for i := range arr {
		if arr[i] < first {
			left = append(left, arr[i])
		}

		if arr[i] > first {
			right = append(right, arr[i])
		}
	}

	res := append(append(quickSort(left), first), quickSort(right)...)

	return res
}

func main() {
	nums := []int{5, 4, 7, 3, 6, 2, 1}
	sortedNums := quickSort(nums)
	fmt.Println(sortedNums)
}
