package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度 on 空间复杂度on
func findTarget(root *TreeNode, k int) bool {
	hash := map[int]int{}
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := hash[k-node.Val]; ok {
			return true
		}
		hash[node.Val] = 1
		return dfs(node.Left) || dfs(node.Right)
	}
	ok := dfs(root)
	return ok
}

func main() {

}
