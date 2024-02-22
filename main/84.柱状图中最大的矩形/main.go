package main

//暴力法 超时 O（N*N）
// func largestRectangleArea(heights []int) int {
// 	max_area := 0
// 	max_length := len(heights)
// 	for i := 0; i < max_length; i++ {
// 		min_h := math.MaxInt32
// 		length := 0

// 		for j := i; j < max_length; j++ {
// 			length += 1
// 			if heights[j] < min_h {
// 				min_h = heights[j]
// 			}
// 			area := min_h * length
// 			if area > max_area {
// 				max_area = area
// 			}
// 			if min_h*max_length < max_area {
// 				break
// 			}
// 		}

// 	}
// 	return max_area
// }

//单调栈，以每根柱子为中心，计算两侧高度高于当前柱子的矩形, 永远保持栈的单调性，因为大于当前柱子的已知柱子会被排除出答案区
func largestRectangleArea(heights []int) (ans int) {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)

	for i := range left {
		left[i] = -1
		right[i] = n
	}

	stk := []int{}

	for i, x := range heights {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= x { //记录左边最近小于当前高度的下标
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	//清空栈
	stk = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	for i, x := range heights {
		ans = max(ans, (right[i]-left[i]-1)*x)
	}
	return
}

//单调栈+常数优化
func largestRectangleArea(heights []int) (ans int) {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)

	for i := range left {
		left[i] = -1
		right[i] = n
	}

	stk := []int{}

	for i, x := range heights {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= x { //记录左边最近小于当前高度的下标
			right[stk[len(stk)-1]] = i
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	//清空栈
	for i, x := range heights {
		ans = max(ans, (right[i]-left[i]-1)*x)
	}
	return
}

func main() {

}
