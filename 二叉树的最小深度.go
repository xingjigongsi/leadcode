/**
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
package leadcode

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	dep_min := math.MaxInt64
	if root.Right != nil {
		dep_min = min(dep_min, minDepth(root.Right))
	}
	if root.Left != nil {
		dep_min = min(dep_min, minDepth(root.Left))
	}
	return dep_min + 1

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//第二种解法

func minDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dep1 := minDepth_1(root.Left)
	dep2 := minDepth_1(root.Right)
	if root.Left == nil || root.Right == nil {
		return dep1 + dep2 + 1
	}
	return min(dep1, dep2) + 1
}

// 第三种解法

func minDepth_2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth_2(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth_2(root.Left) + 1
	}
	return min(minDepth_2(root.Right), minDepth_2(root.Left)) + 1
}
