package main

import "fmt"

/*
前缀和计算
*/
func pivotIndex(nums []int) int {
	n := len(nums)
	prex := make([]int, n+1)
	for i, v := range nums {
		prex[i+1] = prex[i] + v
	}

	for j := 0; j < n; j++ {
		if prex[j] == prex[n]-prex[j]-nums[j] {
			return j
		}
	}

	return -1
}

/*
优化，不必一开始计算所有前缀和，动态计算
*/
func pivotIndex(nums []int) int {
	l, r := 0, 0
	for _, v := range nums {
		r += v
	}

	for i, v := range nums {
		r -= v
		fmt.Println(l, v)
		if l == r {
			return i
		}
		l += v
	}

	return -1
}

/*
进一步优化，只求前缀，只要求2*prex+num == total即可
*/
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1
}
func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
