package main

import (
	"fmt"
	"strconv"
)

//string.Count(fmt.Sprintf("%b", i ), "1")  i转化为二进制格式，并检查1的个数

// O(n* a)做法,先遍历n数，针对每个n进行检测二进制码
// 优化？
func countBits(n int) []int {
	var res = make([]int, n+1)
	res[0] = 0
	for i := 1; i <= n; i++ {
		str := strconv.FormatInt(int64(i), 2)
		fmt.Printf("%s,   %T\n", str, str)

		//遍历acs2编码
		for _, v := range str {
			// fmt.Println(k)
			if v == '1' {
				res[i]++
			}
		}
	}
	return res
}

// Brain Kernighan算法，该（x的二进制表示的最后一个1变成0），那么循环该过程就是x的一比特数， 例10000 & 01111 每次都会去除最后一个1，当x为零时则为1的个数 时间复杂度为0(nlogn)
func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

/*
动态规划 Brain Kernighan算法, 对任意整数x 有x= x&（x-1），该（x的二进制表示的最后一个1变成0），那么循环该过程就是x的一比特数， 例10000 & 01111 每次都会去除最后一个1
（1）最低有效位，每次去掉最低有效位，f[0] = 0, f[1]=1, f[2]=1 当最后一位变为1时，f[i] = f[i-1]+1,那么当我为 我去掉一个1，为我前一个数据，此时已经记录，那么res就可以根据i&i-1时必然小于i，那么前面的数值已经记录
时间复杂度o（n）， 空间复杂度 O（n）

（2）最高有效位，每次找当前i的最高有效位，每次去掉最高有效位则有f[i] = f[i-high_bit] +1， 那么我就可以根据减去最高有效位的下标+1就可以通过动态规划进行计算，因为总有x-high_bit = y x》y
*/

// 最低有效位
func countBits(n int) []int {
	var res = make([]int, n+1)
	res[0] = 0
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

// 最高有效位
func countBits(n int) []int {
	var res = make([]int, n+1)
	res[0] = 0
	high_bit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			high_bit = i
		}
		res[i] = res[i-high_bit] + 1
	}
	return res
}

/*
二进制奇偶性原理，在二进制中 奇数的最后一位必定为1， 偶数的最后一位必定为0， 那么在判断一个数所包含一的个数有以下推论：
① 当前为奇数时，我所拥有的1的个数为 我前一个数的1的数量+1
② 当前为偶数时，我所拥有1的个数为 二进制右移一位的个数，因为偶数最后一位必定为0
*/
func countBits(n int) []int {
	var res = make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i-1] + 1
		}
	}
	return res
}

func main() {
	fmt.Println(countBits(2))
}
