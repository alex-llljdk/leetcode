package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//创建两个队列，存放当前层和下一层，通过不断遍历当前层计算出每层最大值
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := []*TreeNode{root}
	tmp := []*TreeNode{}
	max1 := math.MinInt64
	for len(queue) != 0 {
		if queue[0].Val > max1 {
			max1 = queue[0].Val
		}
		if queue[0].Left != nil {
			tmp = append(tmp, queue[0].Left)
		}
		if queue[0].Right != nil {
			tmp = append(tmp, queue[0].Right)
		}
		queue = queue[1:]
		if len(queue) == 0 {
			res = append(res, max1)
			queue = tmp
			tmp = []*TreeNode{}
			max1 = math.MinInt64
		}
	}
	return res
}

//只用一个序列，每次都用循环取出队列中的所有数值，并记录
func largestValues(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	var q = []*TreeNode{root}
	for len(q) > 0 {
		t := math.MinInt32
		for i := len(q); i > 0; i-- {
			node := q[0]
			q = q[1:]
			t = max(t, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, t)
	}
	return ans
}

func main() {

}
