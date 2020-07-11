package leadcode

import (
	"exam/tool"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *tool.ListNode, l2 *tool.ListNode) *ListNode {
	new_list := &ListNode{}
	temp := new_list
	for {
		if l1 == nil {
			temp.Next = &ListNode{Val: l2.Val}
			temp = temp.Next
			l2 = l2.Next
			break
		}
		if l2 == nil {
			temp.Next = &ListNode{Val: l1.Val}
			temp = temp.Next
			l1 = l1.Next
			break
		}
		if l1.Val <= l2.Val {
			temp.Next = &ListNode{Val: l1.Val}
			temp = temp.Next
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			temp.Next = &ListNode{Val: l2.Val}
			temp = temp.Next
			l2 = l2.Next
		}
	}
	return new_list.Next
}

func showNode(head *ListNode) {
	temp := head.Next
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}
