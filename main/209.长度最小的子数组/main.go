package main

import (
	"fmt"
)

/*
双指针法，滑动窗口
都是正数
指定左右两个指针，当sum<target时，右指针一直往右走， 当sum>=target时，左值针往左走，同时减去前一个值，此时检查sum是否符合条件。不符合则继续移动右指针，直到求出最小长度
*/
// func minSubArrayLen(target int, nums []int) (res int) {
// 	left, right := 0, 0
// 	sum := 0
// 	for ; right < len(nums); right++ {
// 		sum += nums[right]
// 		if sum >= target {
// 			for sum >= target {
// 				sum -= nums[left]
// 				left++
// 			}
// 			if right-left+2 < res || res == 0 {
// 				res = right - left + 2
// 			}

// 		}
// 	}
// 	return
// }

/*
结构优化

	func minSubArrayLen(target int, nums []int) int {
		const inf = 1 << 30
		ans := inf
		s, i := 0, 0
		for j, x := range nums {
			s += x
			for s >= target {
				ans = min(ans, j-i+1)
				s -= nums[i]
				i++
			}
		}
		if ans == inf {
			return 0
		}
		return ans
	}
*/

/*
前缀和+二分查找
*/
func minSubArrayLen(target int, nums []int) (res int) {
	n := len(nums)
	prex := make([]int, n+1)
	res = 1 << 30
	//前缀和
	for i := 1; i < n+1; i++ {
		prex[i] = prex[i-1] + nums[i-1]
	}

	//binary search
	for i := range prex {
		left, right := i, len(prex)-1
		for left <= right {
			mid := left + (right-left)>>1
			fmt.Println(i, mid, left, right)
			if prex[mid]-prex[i] >= target && prex[mid-1]-prex[i] < target {
				res = min(mid-i, res)
				break
			} else if prex[mid]-prex[i] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	if res == 1<<30 {
		return 0
	}
	return res
}

// func minSubArrayLen(s int, nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
// 	ans := math.MaxInt32
// 	sums := make([]int, n+1)
// 	// 为了方便计算，令 size = n + 1
// 	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
// 	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
// 	// 以此类推
// 	for i := 1; i <= n; i++ {
// 		sums[i] = sums[i-1] + nums[i-1]
// 	}
// 	for i := 1; i <= n; i++ {
// 		target := s + sums[i-1]                //预处理前缀和
// 		bound := sort.SearchInts(sums, target) //找到target在sums中的位置，比较前缀和，那么bound就是最合适的前缀和

// 		if bound <= n { //如果该位置在合理位置，比较差值和最终答案
// 			ans = min(ans, bound-(i-1))
// 		}
// 	}
// 	if ans == math.MaxInt32 {
// 		return 0
// 	}
// 	return ans
// }

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
