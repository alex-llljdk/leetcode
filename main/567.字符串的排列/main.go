package main

/*
双指针滑动窗口 O(N+M)
*/
func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	var cnt1, cnt2 [26]int
	for i := 0; i < m; i++ {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := m; i < n; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-m]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//优化，通过中间变量判断是否相等，不必比较两个数组
func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	var cnt [26]int
	diff := 0
	for i := 0; i < m; i++ {
		cnt[s1[i]-'a']--
		cnt[s2[i]-'a']++
	}
	for _, x := range cnt {
		if x != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}

	for i := m; i < n; i++ {
		a, b := s2[i-m]-'a', s2[i]-'a'
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
			return true
		}
	}
	return false
}

/*
双指针法 O(N+M)
*/
func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	//统计字符数为负
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		//遍历检测右指针，如果右指针当前值在s1中有，且数量大于现有值，那么继续右移指针
		//如果右指针当前值在s1中没有或者已经超过s1存在的值，那么我们的左值针就需要右移，直到数量值重新相等
		//当且仅当左右指针的距离==n时，则cnt中负数全部被增加为0，有结果
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}

func main() {

}
