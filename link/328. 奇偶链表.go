package link

type ListNode struct {
	Val int
	Next *ListNode
 }

func oddEvenList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	even := head  //奇数尾部
	odd:= head.Next //偶数尾部
	odd_end:= odd  //偶数头
	for odd!=nil && odd.Next!=nil{
		even.Next = odd.Next
		even = even.Next
		odd.Next = even.Next
		odd = odd.Next
	}
	even.Next = odd_end
	return head
}
