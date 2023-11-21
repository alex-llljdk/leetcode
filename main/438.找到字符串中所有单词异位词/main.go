package main

import "fmt"

func findAnagrams(s string, p string) (res []int) {
	n, m := len(s), len(p)
	cnt := [26]int{}
	diff := 0

	for i := 0; i < m; i++ {
		cnt[p[i]-'a']++
		diff++
	}

	i, j := 0, 0
	for ; i < n; i++ {
		fmt.Println(diff)
		if cnt[s[i]-'a'] > 0 {
			diff--
		}
		cnt[s[i]-'a']--

		for i-j >= m {
			if cnt[s[j]-'a'] <= 0 {
				diff++
			}
			cnt[s[j]-'a']++
			j++
		}

		if diff == 0 {
			res = append(res, j)
		}
	}
	return
}

/*
滑动窗口法 使用两个数组保存,每次进入一个最新数据，出来一个久的数据
*/
func findAnagrams(s string, p string) (res []int) {
	m, n := len(s), len(p)
	if m < n {
		return
	}
	var cnt1, cnt2 [26]int
	for i, ch := range p {
		cnt1[s[i]-'a']++
		cnt2[ch-'a']++
	}
	if cnt1 == cnt2 {
		res = append(res, 0)
	}
	for i := n; i < m; i++ {
		cnt1[s[i]-'a']++
		cnt1[s[i-n]-'a']--
		if cnt1 == cnt2 {
			res = append(res, i-n+1)
		}
	}
	return
}

/*
滑动窗口优化 使用diff进行储存，不必每次比较两个数组，节省时间
*/
func findAnagrams(s string, p string) (res []int) {
	m, n := len(s), len(p)
	if m < n {
		return
	}
	cnt, diff := [26]int{}, 0
	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
		cnt[p[i]-'a']--
	}

	for _, x := range cnt {
		if x != 0 {
			diff++
		}
	}
	if diff == 0 {
		res = append(res, 0)
	}
	for i := n; i < m; i++ {
		a, b := s[i-n]-'a', s[i]-'a'
		if cnt[a] == 0 {
			diff++
		}
		cnt[a]--
		if cnt[a] == 0 {
			diff--
		}

		if cnt[b] == 0 {
			diff++
		}
		cnt[b]++
		if cnt[b] == 0 {
			diff--
		}
		if diff == 0 {
			res = append(res, i-n+1)
		}
	}
	return
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))

}
