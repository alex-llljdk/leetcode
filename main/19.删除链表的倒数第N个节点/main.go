package main

/*
快慢指针法 + 虚拟头节点
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast := dummy
	slow := dummy
	for n > 0 {
		fast = fast.Next
		n -= 1
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {

}
