package leadcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var results []int

func inorderTraversal(root *TreeNode) []int {
	results = make([]int, 0)
	helper(root)
	return results
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left)
	results = append(results, root.Val)
	helper(root.Right)
}

// 迭代的第一种遍历
func inorderTraversal2(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		result = append(result, curr.Val)
		stack = stack[:len(stack)-1]
		curr = curr.Right
	}
	return result
}

// 迭代的第二种遍历
func inorderTraversal3(root *TreeNode) []int {
	result := []int{}
	statck := []interface{}{root}
	for len(statck) != 0 {
		temp := statck[len(statck)-1]
		statck = statck[:len(statck)-1]
		if v, ok := temp.(*TreeNode); ok && v != nil {
			statck = append(statck, v.Right, v.Val, v.Left)
		}
		if v, ok := temp.(int); ok {
			result = append(result, v)
		}
	}
	return result
}

func inorderTraversal4(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	stack := []*TreeNode{}
	curr := root
	for curr != nil || len(stack) != 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			curr = node.Right
		}
	}
	return result
}
