package main

/*
方法一 寻找中点、逆转链表、合并链表
方法二 数组记录
方法三 递归设置全局变量
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	mid := findMid(head)
	tmp := mid.Next
	mid.Next = nil
	tmp = reverseList(tmp)
	head = mergeList(head, tmp)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func findMid(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	for l1 != nil && l2 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

func main() {

}
