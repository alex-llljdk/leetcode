package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归做法,当前节点为0且左右节点为空时，就会去除该节点 每个节点遍历一遍时间On  空间On
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

// 复杂dfs做法，效率低
func pruneTree(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left || right {
			if !left {
				node.Left = nil
			}
			if !right {
				node.Right = nil
			}
			return true
		} else {
			node.Left = nil
			node.Right = nil
		}
		if node.Val == 1 {
			return true
		} else {
			return false
		}
	}
	res := dfs(root)
	if res {
		return root
	} else {
		return nil
	}
}

func main() {

}
