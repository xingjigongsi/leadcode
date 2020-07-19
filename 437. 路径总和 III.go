package leadcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil{
		return 0
	}
	res:=pathHelper(root,sum)
	res+=pathSum(root.Left,sum)
	res+=pathSum(root.Right,sum)
	return res
}

func pathHelper(root *TreeNode, sum int) int{
	if root == nil{
		return 0
	}
	res:= 0
	if root.Val == sum{
		res+=1
	}
	res+=pathHelper(root.Left,sum-root.Val)
	res+=pathHelper(root.Right,sum-root.Val)
	return res
}
