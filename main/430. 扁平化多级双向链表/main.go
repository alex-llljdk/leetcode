package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

/*
dfs,每次都进入子节点，进行操作，通过记录last child节点，连接next节点，利用递归栈, 时间复杂度O（n）
*/
func dfs(node *Node) (last *Node) {
	cur := node
	for cur != nil {
		next := cur.Next
		if cur.Child != nil {
			childLast := dfs(cur.Child) //进入子节点，得到子节点的最后一个节点
			//退出子节点，对子节点进行连接
			next = cur.Next

			cur.Next = cur.Child
			cur.Child.Prev = cur

			//将子节点的最后节点与下一节点进行连接
			if next != nil {
				childLast.Next = next
				next.Prev = childLast
			}

			//子节点置0，最后节点设为当前最后子节点
			cur.Child = nil
			last = childLast
		} else {
			last = cur
		}
		cur = next
	}
	return
}

func flatten(root *Node) *Node {
	dfs(root)
	return root
}

/*
flatten 自递归，前序遍历，由于需要每个节点最高可能访问h次，h递归深度，每次都要检查子层中是否拥有子节点，然后还需要获取子结点中的末尾结点故O（n方），每次都要扫一遍以连接节点的末尾
*/
func flatten(root *Node) *Node {
	dummy := *&Node{}
	dummy.Next = root

	for root != nil {
		if root.Child == nil {
			root = root.Next
		} else {
			tmp := root.Next
			childLast := flatten(root.Child) //返回展平的最后一个子节点

			root.Next = childLast
			childLast.Prev = root
			root.Child = nil

			for root.Next != nil { //获取子节点的末尾节点
				root = root.Next
			}
			root.Next = tmp
			if tmp != nil {
				tmp.Prev = root
			}
			root = tmp
		}
	}
	return dummy.Next
}

/*
迭代法, 直接操作当前层，每次都直接连接子节点和当前层的最后子节点，然后再进行迭代寻找下一个子节点，因为每一层的连接的次序都是固定的，所以可以通过迭代子节点继续宁连接 O（n） o1
*/
func flatten(root *Node) *Node {
	dummy := *&Node{}
	dummy.Next = root
	for root != nil {
		if root.Child != nil {
			tmp := root.Next
			child := root.Child

			root.Next = child
			child.Prev = root
			root.Child = nil

			last := root

			for last.Next != nil {
				last = last.Next
			}
			last.Next = tmp
			if tmp != nil {
				tmp.Prev = last
			}
		}
		root = root.Next
	}
	return dummy.Next
}

func main() {

}
