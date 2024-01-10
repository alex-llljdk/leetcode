package main

import (
	"fmt"
	"math"
)

/*
双指针滑动窗口

思路：

	①长度比较，不符合返回空
	②使用数组记录，统计目标字符串中字符出现次数++，同时统计原字符串中字符出现次数--，这样就可以计算出在等长度下中字符的差异位数，同时将>0的次数记录为diff，当<=0时认为当前记录数相同，或原字符串中
		多出的字符，不计入diff中
	③滑动窗口记录最小覆盖字串，当右指针检测到拥有覆盖子串，左值针不断的取出字符，直到覆盖子串消失，此时拥有一个覆盖子串，并进行长度比较
	时间复杂度O(M+N),空间复杂度O（N）
*/
func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if n > m {
		return ""
	}
	cnt := make([]int, 58)
	for i, v := range t { //记录等长下，数量差异值
		cnt[v-'A']++
		cnt[s[i]-'A']--
	}
	diff := 0
	for i := range cnt { //当cnt>0时表示缺省值
		if cnt[i] > 0 {
			diff++
		}
	}

	if diff == 0 {
		return s[:n]
	}

	minLen := math.MaxInt
	r, l := 0, 0
	for i, j := n, 0; i < m; i++ {
		cnt[s[i]-'A']--
		if cnt[s[i]-'A'] == 0 { //右指针解决缺省值，diff--
			diff--
		}
		for diff == 0 { //此时匹配项中，无缺省值
			cnt[s[j]-'A']++        //左值针退字符
			if cnt[s[j]-'A'] > 0 { //判断缺省值是否重新进入cnt，diff++
				diff++
			}
			if i-j+1 < minLen { //检测当前字符串是否位最小覆盖，记录左右下标
				minLen = i - j + 1
				r = j
				l = i
			}
			j++
		}

	}
	if l != 0 {
		return s[r : l+1]
	}
	return ""
}

func main() {
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}
