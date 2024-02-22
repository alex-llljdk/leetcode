package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func findBottomLeftValue(root *TreeNode) int {
	curHeight := 0
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)
		if height > curHeight {
			curHeight = height
			res = node.Val
		}
	}
	dfs(root, 0)
	return res
}

// 每次取出某一层的数据，观察是否拥有下一层
func findBottomLeftValue(root *TreeNode) int {
	tree := []*TreeNode{root}
	res := 0
	for len(tree) != 0 {
		res = tree[0].Val
		for i := len(tree); i > 0; i-- {
			node := tree[0]
			tree = tree[1:]
			if node.Left != nil {
				tree = append(tree, node.Left)
			}
			if node.Right != nil {
				tree = append(tree, node.Right)
			}
		}
	}
	return res
}

func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{root}
	ans := -1
	for n := len(q); n > 0; n = len(q) {
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if i == 0 {
				ans = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return ans
}

func main() {

}
