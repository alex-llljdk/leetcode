package main

type Node struct {
	Val  int
	Next *Node
}

/*
双指针法，要求找到非递减的序列位置插入，那么就可以分为几种情况
①找到一个位置，A<X<B
②当x为最大或最小时，找到最大值点，将该值插入到最大值和最小值之间
*/
func insert(aNode *Node, x int) *Node {
	node := &Node{Val: x}

	if aNode == nil {
		aNode = &Node{x, nil}
		aNode.Next = aNode
		return aNode
	}

	prev, curr := aNode, aNode.Next

	for curr != aNode {
		if prev.Val <= x && x <= curr.Val || prev.Val > curr.Val && (x >= prev.Val || x < curr.Val) {
			break
		}
		prev, curr = curr, curr.Next
	}
	prev.Next = node
	node.Next = curr
	return aNode
}

func main() {

}
