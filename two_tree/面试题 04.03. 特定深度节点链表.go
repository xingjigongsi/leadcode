package two_tree


type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

type ListNode struct {
	 Val int
	 Next *ListNode
}

// return  []*ListNode{&ListNode{1,nil},&ListNode{2,&ListNode{3,nil}}}
func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil{
		return []*ListNode{}
	}
	queue:=[]*TreeNode{tree}
	result:= []*ListNode{}
	for len(queue)!=0{
		l_queue:= len(queue)
		result_temp:= &ListNode{}
		for i:=0;i<l_queue;i++{
			temp := queue[0]
			queue = queue[1:]
			if temp.Left!=nil{
				queue = append(queue,temp.Left)
			}
			if temp.Right!=nil{
				queue = append(queue,temp.Right)
			}
			tmep:=result_temp
			for tmep.Next!=nil{
				tmep = tmep.Next
			}
			tmep.Next = &ListNode{Val:temp.Val}
		}
		result = append(result,result_temp.Next)
	}
	return result
}