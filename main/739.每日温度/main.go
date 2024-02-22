package main

import "math"

//单调栈，找出每个数左右边离他最近的且比该数大/小的数
//后向遍历
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures)) //创建结果数组
	var stk []int                         //创建栈,存储下标
	for i, t := range temperatures {      //遍历数组
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] < t { //遍历单调栈,判断当前温度是否大于栈顶的温度，栈中数据单调存储
			j := stk[len(stk)-1]
			ans[j] = i - j
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return ans
}

//前向遍历
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var stk []int
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] <= temperatures[i] { //如果栈中有元素，且栈顶元素温度小于等于当前温度，要求找到大于当前温度的下标，出栈所有小于当前温度的数据，表示没有大于当前温度的数据
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 { //如果找到了大于当前温度的数据，那么就是当前数据的下标
			ans[i] = stk[len(stk)-1] - i
		}
		stk = append(stk, i) //该数据入栈，保持栈的单调性，且最近
	}
	return ans
}

//暴力法 时间O(MN) O（M）
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length) //记录答案表
	next := make([]int, 101)   //记录最近出现温度下标表
	for i := 0; i < 101; i++ { //初始设置最近出现下标为无穷大
		next[i] = math.MaxInt32
	}
	for i := length - 1; i >= 0; i-- { //反向遍历
		warmerIndex := math.MaxInt32                  //最小距离出现的下标
		for t := temperatures[i] + 1; t <= 100; t++ { //从当前温度开始向上遍历
			if next[t] < warmerIndex { //如果下标小于记录下表，那么替换为更小的下标，记录出最小的下标即为最近温度
				warmerIndex = next[t]
			}
		}
		if warmerIndex < math.MaxInt32 { //记录差值
			ans[i] = warmerIndex - i
		}
		next[temperatures[i]] = i //最后更新下标
	}
	return ans
}

func main() {

}
