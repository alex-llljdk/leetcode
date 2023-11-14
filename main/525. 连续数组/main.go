package main

import "fmt"

/*
数组中只有0，1数值，那么就可以记录前缀和，通过前缀和数量，但是0，1无法确定0的数值，那么就可以通过把0比作-1，计算sum，得出0，1数量的差值，要求得到最大值，那么我就直接记录初次出现的坐标即可
*/
func findMaxLength(nums []int) (ans int) {
	const zero int = -1
	const one int = 1
	s := 0
	hash := map[int]int{0: -1}
	for i, v := range nums {
		if v == 1 {
			s += one
		} else {
			s += zero
		}
		fmt.Println(hash[s])

		if j, ok := hash[s]; ok {
			ans = max(ans, i-j)
		} else {
			hash[s] = i
		}
	}
	return
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
}
