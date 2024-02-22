package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//队列法
type CBTInserter struct {
	queue []*TreeNode
	head  *TreeNode
}

//计算出最后有空的节点队列
func Constructor(root *TreeNode) CBTInserter {
	CBT := CBTInserter{
		queue: []*TreeNode{},
		head:  root,
	}
	CBT.queue = append(CBT.queue, root)
	for CBT.queue[0].Left != nil {
		CBT.queue = append(CBT.queue, CBT.queue[0].Left)
		if CBT.queue[0].Right == nil {
			break
		}
		CBT.queue = append(CBT.queue, CBT.queue[0].Right)
		CBT.queue = CBT.queue[1:]
	}
	return CBT
}

//通过队列插入左右节点
func (this *CBTInserter) Insert(val int) int {
	node := &TreeNode{Val: val}
	if this.queue[0].Left == nil {
		this.queue[0].Left = node
		this.queue = append(this.queue, node)
		return this.queue[0].Val
	}
	this.queue[0].Right = node
	this.queue = append(this.queue, node)
	res := this.queue[0].Val
	this.queue = this.queue[1:]
	return res
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.head
}

//完全二叉树的性质
type CBTInserter struct {
	tree []*TreeNode
}

//数组完全二叉树，编号的一半即是父节点
func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	tree := []*TreeNode{}
	for len(q) > 0 {
		node := q[0]
		tree = append(tree, node)
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return CBTInserter{tree}
}

//数组长度为从1开始编号，-1为从0开始编号
func (this *CBTInserter) Insert(val int) int {
	pid := (len(this.tree) - 1) >> 1
	node := &TreeNode{Val: val}
	this.tree = append(this.tree, node)
	p := this.tree[pid]
	if p.Left == nil {
		p.Left = node
	} else {
		p.Right = node
	}
	return p.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.tree[0]
}

func main() {
	CBT := Constructor(&TreeNode{Val: 1})
	fmt.Println(CBT.Insert(2))
	fmt.Println(CBT.Insert(3))
	fmt.Println(CBT.Insert(4))
}
