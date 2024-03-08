package main

import "fmt"

func quickSortV3(arr []int, left, right int) {
	if left >= right {
		return
	}
	cur, lo := left+1, left
	for cur <= right {
		if arr[cur] <= arr[left] {
			arr[lo+1], arr[cur] = arr[cur], arr[lo+1]
			lo++
		}
		cur++
	}
	arr[left], arr[lo] = arr[lo], arr[left]
	quickSortV3(arr, left, lo-1)
	quickSortV3(arr, lo+1, right)
}

func main() {
	arry := []int{3, 1, 5, 7, 8, 9, 2, 1, 6}
	quickSortV3(arry, 0, len(arry)-1)
	fmt.Println(arry)
}
