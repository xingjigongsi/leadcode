package leadcode

import "math"

func isValidBST(root *TreeNode) bool {
	return helpers(root, math.MinInt64, math.MaxInt64)
}

func helpers(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	val := root.Val
	if val >= upper || val <= lower {
		return false
	}
	return helpers(root.Left, lower, val) && helpers(root.Right, val, upper)
}

func isValidBST_1(root *TreeNode) bool {
	statck := []*TreeNode{}
	curr := root
	result := []int{}
	for curr != nil || len(statck) != 0 {
		if curr != nil {
			statck = append(statck, curr)
			curr = curr.Left
		} else {
			temp := statck[len(statck)-1]
			if len(result) != 0 {
				kk := result[len(result)-1]
				result = result[:len(result)-1]
				if temp.Val < kk {
					return false
				}
			}
			statck = statck[:len(statck)-1]
			result = append(result, temp.Val)
			curr = temp.Right
		}
	}
	return true
}

var pre = math.MinInt16

func isValidBST_2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBST_2(root.Left) {
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return isValidBST_2(root.Right)
}
