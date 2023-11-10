package main

import (
	"fmt"
	"math"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/* 暴力解法+（快速幂），除法相当于多次减法，那么就可以通过循环减进行计算，但是每次只走一个步长会导致超时，那么在判断时选择当前的最长步长进行计算，减少计算量 */
// func divide(dividend int, divisor int) int {
// 	var sign int
// 	if dividend*divisor < 0 {
// 		sign = -1
// 	} else {
// 		sign = 1
// 	}

// 	a := abs(dividend)
// 	b := abs(divisor)

// 	tot := 0

// 	for a >= b {
// 		cnt := 0
// 		for a >= (b << (cnt + 1)) { //按照幂次选择当前最长可走步数
// 			cnt++
// 		}
// 		tot += 1 << cnt
// 		a -= b << cnt
// 	}

// 	ans := sign * tot
// 	if ans >= math.MinInt32 && ans <= math.MaxInt32 {
// 		return ans
// 	}
// 	return math.MaxInt32
// }

/*
二分查找+快速幂
由于在int中，负数的可表示的最大值上限比正数多一，那么可以考虑使用负数进行整体运算，
那么假设X,Y为负数，Z为正数，有X/Y =Z -->  有 Z*Y >= X > (Z+1) * Y

那么我们的目标就是要找到一个Z，使得 MAX（Z*Y）>=X

*/

// func quickAdd(y, z, x int) bool { // 每一轮add = 2* add， z = z/2，那么x = 2^k * y * z/ 2^k
// 	for result, add := 0, y; z > 0; z >>= 1 {
// 		if z&1 > 0 { //判断奇数还是偶数,当奇数位的时候，需要添加余项，因为这样只会保证y = add* 2^k, 当z为基数时，会缺少一个余项，需要补足，第一个条件就是需要add+result > x 即没有除尽
// 			//需要保证 result+add >=x
// 			if result < x-add {
// 				return false
// 			}
// 			result += add
// 		}
// 		if z != 1 { //这里就是每次将y*2，那么对应的z也要除2  ， 当z==0时，即 z*y >x,返回true
// 			//需要保证add+add>=x
// 			if add < x-add {
// 				return false
// 			}
// 			add += add
// 		}
// 	}
// 	return true
// }

// func divide(dividend int, divisor int) int {
// 	if dividend == math.MinInt32 {
// 		if divisor == 1 {
// 			return math.MinInt32
// 		} else if divisor == -1 {
// 			return math.MaxInt32
// 		}
// 	}
// 	if divisor == math.MinInt32 {
// 		if dividend == math.MinInt32 {
// 			return 1
// 		}
// 		return 0
// 	}
// 	//一般情况，二分查找
// 	//将所有的正数取相反数
// 	rev := false
// 	if dividend > 0 {
// 		dividend = -dividend
// 		rev = !rev
// 	}
// 	if divisor > 0 {
// 		divisor = -divisor
// 		rev = !rev
// 	}

// 	ans := 0
// 	left, right := 1, math.MaxInt32
// 	for left <= right {
// 		mid := left + (right-left)>>1 //注意溢出，不能使用除法
// 		if quickAdd(divisor, mid, dividend) {
// 			ans = mid
// 			if mid == math.MaxInt32 { //注意溢出
// 				break
// 			}
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	if rev {
// 		return -ans
// 	}
// 	return ans
// }

// 类二分查找，使用任何数据都可以使用2^k之和表示，那么我们就可以借助空间换时间的方式进行查找，记录y*2^k进行记录，当x在两个区间之间时，那么k就是最大除数，本质还是加速乘
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return dividend
		} else if divisor == -1 {
			return math.MaxInt32
		}
	}

	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}

	sym := false
	if dividend > 0 {
		dividend = -dividend
		sym = !sym
	}
	if divisor > 0 {
		divisor = -divisor
		sym = !sym
	}

	accumulate := []int{divisor}
	for y := divisor; y >= dividend-y; {
		y += y
		accumulate = append(accumulate, y)
	}

	ans := 0
	for i := len(accumulate) - 1; i >= 0; i-- {
		if accumulate[i] >= dividend {
			ans |= 1 << i
			dividend -= accumulate[i]
		}
	}

	if sym {
		return -ans
	}
	return ans
}

func main() {
	fmt.Println(divide(math.MaxInt32, 2))
	fmt.Println(1 & 2)
}
