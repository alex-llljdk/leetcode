package main

import "sort"

//计数法 时间复杂度O（n+n）  空间复杂度O（E）
func isAnagram(s string, t string) bool {
	numList := [26]int{}
	diff := 0

	for i := range s {
		if numList[s[i]-'a'] == 0 {
			diff += 1
		}
		numList[s[i]-'a'] += 1
	}

	for i := range t {
		numList[t[i]-'a'] -= 1
		if numList[t[i]-'a'] < 0 {
			return false
		}

		if numList[t[i]-'a'] == 0 {
			diff -= 1
		}
	}
	if diff == 0 {
		return true
	}
	return false
}

//哈希表法
func isAnagram(s string, t string) bool {
	hash := map[rune]int{}
	diff := 0

	for _, ch := range s {
		_, ok := hash[ch-'a']
		if !ok {
			diff += 1
		}
		hash[ch-'a'] += 1
	}

	for _, ch := range t {
		hash[ch-'a'] -= 1
		if hash[ch-'a'] < 0 {
			return false
		}
		if hash[ch-'a'] == 0 {
			diff -= 1
		}
	}
	if diff == 0 {
		return true
	}
	return false
}

//排序法
func isAnagram(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

func main() {

}
