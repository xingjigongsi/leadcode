package leadcode

/**
题目:
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	statck := []*TreeNode{root}
	var index int
	for i := 1; i < len(preorder); i++ {
		pre_index := preorder[i]
		node := statck[len(statck)-1]
		if node.Val != inorder[pre_index] {
			node.Left = &TreeNode{Val: pre_index}
			statck = append(statck, node.Left)
		} else {
			for len(statck) != 0 && statck[len(statck)-1].Val == inorder[index] {
				node = statck[len(statck)-1]
				index++
				statck = statck[:len(statck)-1]
			}
			node.Right = &TreeNode{Val: pre_index}
			statck = append(statck, node.Right)
		}
	}
	return root
}
