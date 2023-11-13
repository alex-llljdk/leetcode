package main

/*
枚举法
*/
func subarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}

/*
哈希表+前缀和，因为当前数组有正数有负数，故不能使用双指针，时间O(N)，空间O(N)
*/
// func subarraySum(nums []int, k int) int {
// 	hash := map[int]int{0: 1} //哈希记录前缀出现次数
// 	res := 0
// 	s := 0 //记录前缀和
// 	for _, x := range nums {
// 		s += x
// 		res += hash[s-k]
// 		hash[s]++
// 	}
// 	return res
// }

func main() {

}
