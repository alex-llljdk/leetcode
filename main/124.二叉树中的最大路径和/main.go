package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法 时间复杂度O（n） 空间复杂度O（n）
func maxPathSum(root *TreeNode) int {
	ans := -1001
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := max(0, dfs(root.Left))
		r := max(0, dfs(node.Right))

		ans = max(ans, l+r+root.Val)
		return max(l, r) + root.Val
	}
	dfs(root)
	return ans
}
