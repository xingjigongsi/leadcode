package leadcode


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var leve_result [][]int
func levelOrder_11(root *TreeNode) [][]int {
	leve_result = make([][]int,0)
	helper_level(root,0)
	return leve_result
}

func helper_level(root *TreeNode,level int){
	if root == nil{
		return
	}
	if level == len(leve_result){
		result = append(result,[]int{})
	}
	result[level] = append(result[level],root.Val)
	helper_level(root.Left,level+1)
	helper_level(root.Right,level+1)
}

// 迭代法
func levelOrder_12(root *TreeNode) [][]int{
	result:= make([][]int,0)
	q:=[]*TreeNode{root}
	for len(q)!=0{
		q_size:=len(q)
		temp:=[]int{}
		for i:=0;i<q_size;i++{
			q_temp:=q[0]
			q =q[1:]
			temp=append(temp,q_temp.Val)
			if q_temp.Left!=nil{
				q = append(q,q_temp.Left)
			}
			if q_temp.Right!=nil{
				q = append(q,q_temp.Right)
			}
		}
		result = append(result,temp)
	}
	return result
}