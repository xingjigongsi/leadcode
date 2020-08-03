package link


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	pre:=head
	var cur *ListNode
	for pre!=nil{
		t:=pre.Next
		pre.Next = cur
		cur = pre
		pre = t
	}
	return cur
}

func reverseList_1(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	ret:=reverseList_1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}