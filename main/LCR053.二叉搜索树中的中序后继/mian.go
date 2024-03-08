package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var found bool
	var result *TreeNode
	var dfs func(root, p *TreeNode)
	dfs = func(root, p *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left, p)
		if found {
			result = root
			found = false
		}
		if p == root {
			found = true
		}
		dfs(root.Right, p)
	}
	dfs(root, p)
	return result
}

// 迭代法进行中序遍历
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	var prev, cur *TreeNode = nil, root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev == p {
			return cur
		}
		prev = cur
		cur = cur.Right
	}
	return nil
}

// 二叉搜索树，要么在该节点的右子树中的最小值，要么要找到其祖先节点
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode
	if p.Right != nil {
		successor = p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}
	node := root
	for node != nil {
		if node.Val > p.Val {
			successor = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return successor
}

func main() {

}
