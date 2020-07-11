package leadcode

import "math"

/**
您需要在二叉树的每一行中找到最大的值。

示例：

输入:

          1
         / \
        3   2
       / \   \
      5   3   9

输出: [1, 3, 9]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	result := []int{}
	p := []*TreeNode{root}
	for len(p) != 0 {
		var len_p = len(p)
		min := math.MinInt64
		for i := 0; i < len_p; i++ {
			temp := p[0]
			p = p[1:]
			min = max(temp.Val, min)
			if temp.Left != nil {
				p = append(p, temp.Left)
			}
			if temp.Right != nil {
				p = append(p, temp.Right)
			}
		}
		result = append(result, min)
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var largest_result = []int{}

func largestValues_1(root *TreeNode) []int {
	largest_result = make([]int, 0)
	largestValues_helper(root, 0)
	return largest_result
}

func largestValues_helper(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if len(largest_result) == level {
		largest_result = append(largest_result, math.MinInt64)
	}
	largest_result[level] = max(largest_result[level], root.Val)
	if root.Left != nil {
		largestValues_helper(root.Left, level+1)
	}
	if root.Right != nil {
		largestValues_helper(root.Right, level+1)
	}
}
