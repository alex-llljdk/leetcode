package main

import (
	"fmt"
	"sort"
)

/*
排序+双指针, 排序的目的是为了跳过重复值，已知重复值的 i+j+k已经计算那么就没有必要计算重复值
时间复杂度O（n*n） 空间复杂度O（1）
*/
func threeSum(nums []int) (res [][]int) {

	l := len(nums)
	// sort.SliceIsSorted(nums, func(i, j int) bool {
	// 	return nums[i] < nums[j]
	// })
	sort.Ints(nums)

	for i := 0; i < l-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		head, tail := i+1, l-1

		for head < tail {
			sum := nums[i] + nums[head] + nums[tail]
			if sum < 0 {
				head++
			} else if sum > 0 {
				tail--
			} else {
				res = append(res, []int{nums[i], nums[head], nums[tail]})
				head++
				tail--
				for head < tail && nums[head] == nums[head-1] {
					head++
				}
				for head < tail && nums[tail] == nums[tail+1] {
					tail--
				}
			}
		}

	}
	return
}
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
