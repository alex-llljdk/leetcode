/* 
方法，找到回文串的中点，将后半部分进行逆转，然后开始进行比较即可
*/
func isPalindrome(head *ListNode) bool {

	if head.Next=nil{
		return true
	}

	findMid := func(h *ListNode) *ListNode {
		slow, fast := h, h
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}
	reverse := func(h *ListNode) *ListNode {
		var prev *ListNode
		cur := h
		for cur != nil {
			tmp := cur.Next
			cur.Next = prev
			prev = cur
			cur = tmp
		}
		return prev
	}


	mid:= findMid(head)
	tmp:= reverse(mid)

	for tmp!=nil{
		fmt.Println(tmp.Val)
	}

	for tmp!=nil && head!=nil{
		if tmp.Val != head.Val{
			return false
		}
		tmp = tmp.Next
		head = head.Next
	}

	return true
}


/* 
格式----优化
*/
func isPalindrome(head *ListNode) bool {
	if head==nil|| head.Next==nil{
		return true
	}
	slow,fast:= head,head.Next
	for fast!=nil && fast.Next!=nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	cur:=slow.Next
	for cur!=nil{
		t :=cur.Next
		cur.Next = prev
		prev = cur
		cur = t
	}
	for prev!=nil{
		if prev.Val != head.Val{
			return false
		}
		prev = prev.Next
		head = head.Next
	}
	return true
}