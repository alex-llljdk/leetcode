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

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
