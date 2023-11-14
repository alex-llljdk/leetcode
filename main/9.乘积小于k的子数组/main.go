package main

import (
	"fmt"
	"math"
	"sort"
)

/*
双指针滑动窗口
*/
func numSubarrayProductLessThanK(nums []int, k int) (res int) {
	s := 1
	i := 0
	for j, x := range nums {
		s *= x
		for i <= j && s >= k {
			s /= nums[i]
			i++
		}
		res += j - i + 1
	}
	return res
}

/*
二分查找,原理所有都是正数，那么前缀和为递增
*/
func numSubarrayProductLessThanK(nums []int, k int) (res int) {
	n := len(nums)
	prex := make([]float64, n+1)
	for i := 1; i < n+1; i++ {
		prex[i] = prex[i-1] + math.Log(float64(nums[i-1]))
	}

	logK := math.Log(float64(k))

	for i := 1; i <= n; i++ {
		j := sort.SearchFloat64s(prex[:i], prex[i]-logK+1e-10)
		res += i - j
	}
	return res
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
