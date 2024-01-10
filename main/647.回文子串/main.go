package main

/*
暴力双指针，固定尾部指针，检测每个头指针开始到尾指针是否符合回文串O(n*n)
*/
func countSubstrings(s string) int {
	m := len(s)
	res := 0

	check := func(i, j int) bool {
		for i < j {
			if s[i] == s[j] {
				i++
				j--
			} else {
				return false
			}
		}
		return true
	}

	for end := 0; end < m; end++ {
		res++
		for i := 0; i < end; i++ {
			if check(i, end) {
				res++
			}
		}
	}
	return res
}

/*
中心拓展法，由于中心拓展能够更快的找到该串是否为回文串，故时间上有优化
*/
func countSubstrings(s string) int {
	m := len(s)
	res := 0

	f := func(i, j int) (cnt int) {
		for ; i >= 0 && j < m && s[i] == s[j]; i, j = i-1, j+1 {
			cnt++
		}
		return
	}

	for i := range s {
		res += f(i, i)
		res += f(i, i+1)
	}
	return res
}

/*
manacher 算法 线性求解最长回文子串算法，
核心点在于减少遍历次数， 记录最右回文边界，中心位置。 imax,rmax
我们要做的是遍历中心位置，当中心位置在最右回文串里，那么我们就可以通过对称，将过去的回文串复制过来，减少该中心的遍历。
需要在每个字符之间加入#号，保证奇数位的字符串
O（n） O（n）
*/
func countSubstrings(s string) int {
	n := len(s)
	t := "$#"
	for i := 0; i < n; i++ { //添加#号，保证奇数位字符串 #1#2#2#
		t += string(s[i]) + "#"
	}
	n = len(t)
	t += "!" //结束符

	f := make([]int, n) //最长回文半径记录表
	iMax, rMax, ans := 0, 0, 0
	for i := 1; i < n; i++ {
		// 初始化 f[i]
		if i <= rMax { //如果中心点i在最右边界内，那么就记录f[i]目前的最大半径，要么是中心对称，要么是等于最右边界到i的长度
			f[i] = min(rMax-i+1, f[2*iMax-i])
		} else { //如果超出了最右边界，那么此时半径长度为1
			f[i] = 1
		}
		// 中心拓展，通过半径来计算两边的对称值，进行自增
		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}
		// 动态维护 iMax 和 rMax  如果此时最右边界超出当前边界，那么重新赋值中心点和最右边界
		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}
		// 统计答案, 当前贡献为 (f[i] - 1) / 2 上取整 通过每个中心点的半径长度就可以算出，多少个回文子串
		ans += f[i] / 2
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

}
