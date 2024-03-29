package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前缀和+ 哈希表 + 递归
func pathSum(root *TreeNode, targetSum int) int {
	cnt := map[int]int{0: 1}
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, s int) int {
		if node == nil {
			return 0
		}
		s += node.Val
		ans := cnt[s-targetSum]
		cnt[s]++
		ans += dfs(node.Left, s) + dfs(node.Right, s)
		cnt[s]--
		return ans
	}
	return dfs(root, 0)
}

func main() {

}
