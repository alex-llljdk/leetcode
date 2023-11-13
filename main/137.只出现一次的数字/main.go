package main

import (
	"fmt"
)

/*
哈希表法， 将数据存放在哈希表中，根据value值取出为1的key, 时间复杂度为O(N),空间复杂度为O[N]
*/
// func singleNumber(nums []int) int {
// 	m := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		m[nums[i]] += 1
// 	}
// 	for k, v := range m {
// 		if v == 1 {
// 			return k
// 		}
// 	}
// 	return 0
// }

/*
位运算方法，依次考虑二进制编码	统计所有数字中每个位上出现1的数，然后对3取模，如果某一位上出现1的个数不能被3整除，说明只出现一次的数字在该为上是1，否则是0,对每个位置进行遍历
时间复杂度O（N*logm） n是数组长度，M是数组中元素的最大值 空间复杂度O（1）
*/
func singleNumber(nums []int) int {
	var res int32
	for i := 0; i < 32; i++ {
		cnt := 0
		for _, x := range nums {
			cnt += x >> i & 1
		}
		// if cnt%3 != 0 {
		// 	res += int32(math.Pow(2, float64(i)))
		// }
		cnt %= 3
		res |= int32(cnt) << i

	}
	return int(res)
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}
