package leadcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	stack := []interface{}{root}
	result := []int{}
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if v, ok := curr.(*TreeNode); ok && v != nil {
			stack = append(stack, v.Right, v.Left, v.Val)
		}
		if v, ok := curr.(int); ok {
			result = append(result, v)
		}
	}
	return result
}

var result_pro []int

func preorderTraversa1_1(root *TreeNode) []int {
	result_pro = make([]int, 0)
	preorder_helpers(root)
	return result_pro
}

func preorder_helpers(root *TreeNode) {
	if root == nil {
		return
	}
	result_pro = append(result_pro, root.Val)
	preorder_helpers(root.Left)
	preorder_helpers(root.Right)
}

func preorderTraversal_2(root *TreeNode) []int {
	var result_pro []int
	var stack = []*TreeNode{root}
	if root == nil {
		return []int{}
	}
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		result_pro = append(result_pro, curr.Val)

	}
	return result_pro
}

// 第三种写法
func preorderTraversal_3(root *TreeNode) []int {
	result_pro := []int{}
	stack := []*TreeNode{}
	if root != nil {
		result_pro = append(result_pro, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		root = root.Left
		if root == nil && len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return result_pro
}

// 第四种写法
func preorderTraversal_4(root *TreeNode) []int {
	result_pro := []int{}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			result_pro = append(result_pro, node.Val)
			stack = append(stack, node.Right, node.Left)
		}
	}
	return result_pro
}
