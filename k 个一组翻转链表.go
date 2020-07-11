package leadcode

import (
	"fmt"
)

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func addNode(head *ListNode2, Node *ListNode2) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = Node
}

func (ListNode *ListNode2) showNode(head *ListNode2) {
	temp := head.Next
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}

func (this *ListNode2) reverseKGroup(head *ListNode2, k int) *ListNode2 {
	var dummy = &ListNode2{}
	dummy.Next = head // 经典写法
	pre := dummy
	end := dummy
	for {
		if end.Next == nil {
			break
		}
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		Next := end.Next
		end.Next = nil
		pre.Next = this.swapNode(start)
		start.Next = Next
		pre = start
		end = pre
	}
	return dummy
}

func (this *ListNode2) swapNode(start *ListNode2) *ListNode2 {
	curr := start
	var pre *ListNode2
	for {
		if curr == nil {
			break
		}
		temp := curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}
	return pre
}
