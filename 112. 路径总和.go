package leadcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSum_helper(root,sum)
}

func hasPathSum_helper(root *TreeNode, sum int) bool{
	if root == nil{
		return false
	}
	if root.Left == nil && root.Right==nil{
		return root.Val == sum
	}
	return hasPathSum_helper(root.Left,sum-root.Val) ||
		hasPathSum_helper(root.Right,sum-root.Val)
}