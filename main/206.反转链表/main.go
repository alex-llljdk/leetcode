package main

/*
迭代法, 三指针法 或 头插法
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1, p2 := head, head.Next
	p1.Next = nil
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}

/*
递归法
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {

}
