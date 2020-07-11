/**
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

  Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
   Right *TreeNode
   }
*/
package leadcode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max_1(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max_1(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxDepth_1(root *TreeNode) int {
	statck := []map[*TreeNode]int{}
	if root != nil {
		statck = append(statck, map[*TreeNode]int{root: 1})
	}
	depth := 0
	for len(statck) != 0 {
		current := statck[0]
		statck = statck[1:]
		var node *TreeNode
		var curr_dep int
		for key, value := range current {
			node = key
			curr_dep = value
		}
		if node != nil {
			depth = max(depth, curr_dep)
			statck = append(statck, map[*TreeNode]int{node.Left: curr_dep + 1})
			statck = append(statck, map[*TreeNode]int{node.Right: curr_dep + 1})

		}
	}
	return depth
}
