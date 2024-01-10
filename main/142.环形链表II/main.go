package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
哈希表法，如果在循环中碰到相同的节点，那么返回该节点O(N) O(N)
*/
func detectCycle(head *ListNode) *ListNode {
	hash := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := hash[head]; ok {
			return head
		}
		hash[head] = struct{}{}
		head = head.Next
	}
	return nil
}

/*
快慢指针法，设置快指针和慢指针，快指针每次走两步，慢指针每次走一步，
那么当两个指针在环中相遇时有走过的距离，fast = 2slow
设a,b,c分别为入环点距离，相遇点距离，出环点距离
有a+b + N(b+c) = 2*(a+b)，快指针已经走了n圈
有a = c+(n-1)(b+c),那么此时就有入环点距离 = c+N倍圈长
*/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	hasCycle := false
	for !hasCycle && fast != nil && fast.Next != nil { //快慢指针记录相遇点
		slow, fast = slow.Next, fast.Next.Next
		hasCycle = slow == fast
	}
	if !hasCycle {
		return nil
	}
	p := head
	for p != slow { //此时slow在相遇点，只需要走c+（n-1）圈长就可以在相遇点跟ptr相遇
		p, slow = p.Next, slow.Next
	}
	return p
}

func main() {

}
