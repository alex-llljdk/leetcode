package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 反序中序遍历 时间复杂度O（n） 空间复杂度O（n）
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func main() {

}
