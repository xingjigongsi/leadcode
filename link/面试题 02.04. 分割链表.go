package link

func partition(head *ListNode, x int) *ListNode {
	big_head:= &ListNode{}
	small_head:= &ListNode{}
	big_temp:=big_head
	small_temp:= small_head
	head_temp:=head
	for head_temp!=nil{
		if head_temp.Val<x{
			small_temp.Next = &ListNode{Val:head_temp.Val}
			small_temp = small_temp.Next
		}else{
			big_temp.Next = &ListNode{Val:head_temp.Val}
			big_temp = big_temp.Next
		}
		head_temp = head_temp.Next
	}
	small_temp.Next = big_head.Next
	big_temp.Next = nil
	return small_head.Next
}
