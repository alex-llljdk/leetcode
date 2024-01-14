package main

import (
	"sort"
	"strings"
)

//哈希表法，优先将字符串进行排序，这样就可以得到异位词的相同形式
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}

	hash := map[string][]string{}
	for _, val := range strs {
		t := []byte(val)
		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})
		
		k := string(t)
		hash[k] = append(hash[k],val)
	}
	
	ans := [][]string{}
	for key:=range hash{
		ans = append(ans, hash[key])
	}
}

//计数法
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}

	hash := map[[26]int][]string{}
	for _, val := range strs {
		cnt:=[26]int{}
		for _,char:=range val{
			cnt[char-'a']++
		}
		hash[cnt] = append(hash[cnt], val)	
	}
	
	ans := [][]string{}
	for key:=range hash{
		ans = append(ans, hash[key])
	}

	return ans
}



func check(

))
