package leadcode

// 非递归写法
func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []interface{}{root}
	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if v, ok := temp.(*TreeNode); ok && v != nil {
			stack = append(stack, v.Val, v.Right, v.Left)
		}
		if v, ok := temp.(int); ok {
			result = append(result, v)
		}
	}
	return result
}

/**
*  最开始访问的节点是最后访问的，所以后压榨
 */
func postorderTraversal_1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		result = append([]int{node.Val}, result...)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return result
}

func postorderTraversal_2(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	if root == nil {
		return []int{}
	}
	curr := root
	for curr != nil || len(stack) != 0 {

		if curr != nil {
			stack = append(stack, curr)
			result = append([]int{curr.Val}, result...)
			curr = curr.Right
		} else {
			node := stack[len(stack)-1]
			curr = node.Left
			stack = stack[:len(stack)-1]
		}
	}
	return result
}
