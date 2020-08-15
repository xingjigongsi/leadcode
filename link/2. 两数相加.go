package link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	t1,t2,v4:=l1,l2,0
	l3_temp:= l3.Next
	for t1!=nil && t2!=nil{
		va1,va2:=0,0
		if t1!=nil{
			va1 = t1.Val
		}
		if t2!=nil{
			va2 = t2.Val
		}
		temp:=va2+va1+v4
		v4 = 0
		var node *ListNode
		if temp<10{
			node = &ListNode{Val:temp}
		}else{
			node = &ListNode{Val:temp-10}
			v4 = 1
		}
		for l3_temp!=nil{
			l3_temp = l3_temp.Next
		}
		l3_temp = node
		t1 = t1.Next
		t2 = t2.Next
	}
	if v4!=0{
		for l3_temp!=nil{
			l3_temp = &ListNode{Val:v4}
		}
	}
	return l3.Next
}