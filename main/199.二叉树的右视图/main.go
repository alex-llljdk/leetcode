package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
func rightSideView(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		if len(res) < height {
			res = append(res, node.Val)
		}
		dfs(node.Right, height+1)
		dfs(node.Left, height+1)
	}
	dfs(root, 1)
	return res
}

// 层序遍历BFS
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	tree := []*TreeNode{root}
	res := []int{}
	for len(tree) > 0 {
		for i := len(tree); i > 0; i-- {
			node := tree[0]
			if i == 1 {
				res = append(res, tree[0].Val)
			}
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

func main() {

}
