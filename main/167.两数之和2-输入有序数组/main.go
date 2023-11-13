package main

import "fmt"

/*
二分法  时间复杂度O(nlogn)  空间复杂度O(1)
*/
// func twoSum(numbers []int, target int) []int {

// 	for i := 0; i <= len(numbers)-1; i++ {

// 		var left, right int = i + 1, len(numbers) - 1
// 		for left <= right {
// 			mid := left + (right-left)/2

// 			if numbers[i]+numbers[mid] < target {
// 				left = mid + 1
// 			} else if numbers[i]+numbers[mid] > target {
// 				right = mid - 1
// 			} else {
// 				return []int{i + 1, mid + 1}
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }

/*
双指针法	时间复杂度O(n)  空间复杂度O(1)
*/
func twoSum(numbers []int, target int) []int {
	head, tail := 0, len(numbers)-1

	for head < tail {
		sum := numbers[head] + numbers[tail]
		if sum < target {
			head++
		} else if sum > target {
			tail--
		} else {
			return []int{head + 1, tail + 1}
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
