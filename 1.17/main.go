package main

import "fmt"

func binarySearch(arr []int, target int) int {

	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 4, 5, 7}
	res := binarySearch(nums, 5)
	fmt.Println(res)
}
