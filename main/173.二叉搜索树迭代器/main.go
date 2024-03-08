package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// type BSTIterator struct {
// 	iter  []int
// 	point int
// }

// 扁平化方法 时间复杂度On 空间复杂度On
// func Constructor(root *TreeNode) BSTIterator {

// 	bstiter := BSTIterator{point: -1}
// 	var dfs func(*TreeNode)
// 	dfs = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		dfs(node.Left)
// 		bstiter.iter = append(bstiter.iter, node.Val)
// 		dfs(node.Right)
// 	}
// 	dfs(root)
// 	return bstiter
// }

// func (this *BSTIterator) Next() int {
// 	this.point += 1
// 	return this.iter[this.point]
// }

// func (this *BSTIterator) HasNext() bool {
// 	if this.point == len(this.iter)-1 {
// 		return false
// 	}
// 	return true
// }

// 单调栈方法
type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{cur: root}
}

func (this *BSTIterator) Next() int {
	for node := this.cur; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	this.cur, this.stack = this.stack[len(this.stack)-1], this.stack[:len(this.stack)-1]
	val := this.cur.Val
	this.cur = this.cur.Right
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.stack) > 0
}

func main() {

}
