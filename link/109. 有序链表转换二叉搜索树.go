package link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 [-10, -3, 0, 5, 9]
*/
var heads *ListNode
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil{
		return nil
	}
	heads = head
	r:=getSize(head)
	return sorted_helper(0,r-1)
}

func sorted_helper(l int,r int) *TreeNode{
	if l>r{
		return nil
	}
	m:= (l+r)/2
	left:= sorted_helper(l,m-1)
	node := &TreeNode{Val:heads.Val}
	node.Left = left
	heads = heads.Next
	node.Right = sorted_helper(m+1,r)
	return node
}

func getSize(head *ListNode) int{
	size:=0
	temp:= head
	for temp!=nil{
		size++
		temp = temp.Next
	}
	return size
}


// 數組
var data []int
func sortedListToBST_1(head *ListNode) *TreeNode{
	if head==nil{
		return nil
	}
	data = make([]int,0)
	temp:=head
	for temp!=nil{
		data = append(data,temp.Val)
		temp = temp.Next
	}

	return sortedListToBST_helper(0,len(data)-1)
}

func sortedListToBST_helper(l,r int) *TreeNode{
	if l>r{
		return nil
	}
	mid:=(l+r)/2
	node:= &TreeNode{Val:data[mid]}
	if l==mid{
		return node
	}
	node.Left = sortedListToBST_helper(r,mid-1)
	node.Right = sortedListToBST_helper(mid+1,r)
	return node
}



