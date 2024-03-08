package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用栈进行中序遍历，开始时遍历左子树所有元素入栈，取出最后一个元素，判断头尾，进行操作，再检查右子树，中序遍历
func increasingBST(root *TreeNode) *TreeNode {

	var head, tail *TreeNode
	stack := make([]*TreeNode, 0)
	cur := root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1] //退栈
		stack = stack[:len(stack)-1]
		if head == nil {
			head = cur
		} else {
			tail.Right = cur
		}
		tail = cur
		cur.Left = nil
		cur = cur.Right
	}
	return head
}

// dfs中序遍历
func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	cur := dummy
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		root.Left = nil
		cur.Right = root
		cur = root
		dfs(root.Right)
	}
	dfs(root)
	return dummy.Right
}

func main() {

}
