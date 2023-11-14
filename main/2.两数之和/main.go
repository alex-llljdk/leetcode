package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
模拟法，当没有完成两个链表时，同时遍历，每次取出l1, l2的值，将两值相加，直至两个链表计算完毕，最后再次计算最高位进位
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var carry int = 0
	var tail *ListNode
	for l1 != nil || l2 != nil { //迭代两个链表
		n1, n2 := 0, 0
		if l1 != nil { //记录当前node1
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil { //记录node2
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry      //计算累加值
		sum, carry = sum%10, sum/10 //计算个位、进位
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

/*
递归写法,改写法以l1为base节点，每次根据l1节点去调整下一个节点，终止条件为判断（l1,l2）是否为空，前一位carry是否进位
每次判断采用后序递归，先处理，再递归
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	return add(l1, l2, 0)
}

func add(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry != 0 {
			return &ListNode{Val: 1}
		}
		return nil
	} //终止条件
	if l1 == nil {
		l1, l2 = l2, l1
	} //保证l1非空
	carry += l1.Val
	if l2 != nil {
		carry += l2.Val
		l2 = l2.Next
	}
	l1.Val = carry % 10
	l1.Next = add(l1.Next, l2, carry/10)
	return l1
}

func main() {
	result := addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}})
	for ; result != nil; result = result.Next {
		fmt.Println(result.Val)
	}

}
