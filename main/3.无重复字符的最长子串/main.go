package main

import (
	"fmt"
)

/*
双指针法
*/
func lengthOfLongestSubstring(s string) (res int) {
	hash := make(map[byte]int)
	j := 0
	for i := 0; i < len(s); i++ {
		hash[s[i]-'a']++
		for hash[s[i]-'a'] > 1 {
			hash[s[j]-'a']--
			j++
		}
		if i-j+1 > res {
			res = i - j + 1
		}
	}
	return
}

/*
优化 细节 使用bool作为存储空间，因为每个字符只需要记录一次
*/
func lengthOfLongestSubstring(s string) (ans int) {
	ss := make([]bool, 128)
	j := 0
	for i, c := range s {
		for ss[c] {
			ss[s[j]] = false
			j++
		}
		ss[c] = true
		ans = max(ans, i-j+1)
	}
	return
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
