package leadcode

type ListNode1 struct {
	num  int
	next *ListNode1
}

func (listNode *ListNode1) addNode(headnode *ListNode1, node *ListNode1) {
	temp := headnode
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
}

//反转  prev 迭代;5
func (this *ListNode1) reverseList(head *ListNode1) *ListNode1 {
	var prev *ListNode1
	curr := head
	for curr != nil {
		nexTemp := curr.next
		curr.next = prev
		prev = curr
		curr = nexTemp

	}
	return prev
}

func (this *ListNode1) swapPairs(head *ListNode1) *ListNode1 {
	var dummy = head
	var current = dummy
	for current.next != nil && current.next.next != nil {
		first := current.next       //1          3
		second := current.next.next // 2    4
		first.next = second.next    // 1->3   3->nil
		current.next = second       // 0 ->2     1->4
		current.next.next = first   // 2->1   3->3
		current = current.next.next // current = 1
	}
	return dummy.next
}

func hasCycle(head *ListNode1) bool {
	curr := head
	data := make(map[ListNode1]*ListNode1)
	result := false
	for curr != nil {
		data[*curr] = curr
		if _, ok := data[*curr]; ok && curr != nil {
			result = true
		}
		curr = curr.next
	}
	return result
}
