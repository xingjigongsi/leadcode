package leadcode


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var OrderBottom [][]int
func levelOrderBottom(root *TreeNode) [][]int {
	OrderBottom = make([][]int,0)
	levelOrder_helper(root,0)
	//for i:=0;i<len(OrderBottom)/2;i++{
	//	OrderBottom[i],OrderBottom[len(OrderBottom)-1-i]=OrderBottom[len(OrderBottom)-1-i],OrderBottom[i]
	//}
	return OrderBottom
}

func levelOrder_helper(root *TreeNode,level int) {
	if root == nil{
		return
	}
	if len(OrderBottom) == level{
		OrderBottom = append(OrderBottom,[]int{})
	}
	levelOrder_helper(root.Left,level+1)
	levelOrder_helper(root.Right,level+1)
	key:= len(OrderBottom)-level-1
	OrderBottom[key] = append(OrderBottom[key],root.Val)
}


/**
 迭代法
 */
func levelOrderBottom_k(root *TreeNode) [][]int {
	result:=make([][]int,0)
	if root==nil{
		return result
	}
	p:=[]*TreeNode{root}
	for len(p)!=0{
		item:=[][]int{{}}
		p_len:=len(p)
		for i:=0;i<p_len;i++{
			p_item:=p[0]
			p = p[1:]
			item[0] =append(item[0],p_item.Val)
			if p_item.Left!=nil{
				p = append(p,p_item.Left)
			}
			if p_item.Right!=nil{
				p = append(p,p_item.Right)
			}
		}
		result = append(item,result...)
	}
	return result
}