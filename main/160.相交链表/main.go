package main

/*
哈希表法，先遍历A链表，再遍历B链表进行查找
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hash := make(map[*ListNode]bool)
	for headA != nil {
		hash[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if hash[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

/*
双指针法，如何判断有无公共交点，那么我们就要保证我们前进的步骤时一致的，
当我们同时遍历两个链表时，由于两个链表的长度不一致，导致判断困难，
那么当使用两个指针分别走完两个链表，并交换链表，那么就可以保持步伐一致
即A+B+公共 = B+A+公共，那么就会在公共节点相遇，如果没有公共节点，那么就返回nil
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}

	}
	return p1
}

func main() {

}
