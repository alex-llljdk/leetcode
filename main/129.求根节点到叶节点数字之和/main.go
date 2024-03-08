package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
func sumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, v int) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return v*10 + node.Val
		}
		left := dfs(node.Left, v*10+node.Val)
		right := dfs(node.Right, v*10+node.Val)
		return left + right
	}
	return dfs(root, 0)
}

var result int = 0

func sumNumbers(root *TreeNode) int {
	defer func() {
		result = 0
	}()
	getSum(root, 0)
	return result
}
func getSum(root *TreeNode, father int) {
	if root == nil {
		return
	}
	//说明该节点是叶子结点
	if root.Left == nil && root.Right == nil {
		result += father*10 + root.Val
		return
	}
	getSum(root.Left, father*10+root.Val)
	getSum(root.Right, father*10+root.Val)
}
