package leadcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
层序遍历
 **/
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		queue = p
	}
	return ret
}

var leavel_result [][]int

func levelOrder_1(root *TreeNode) [][]int {
	leavel_result = make([][]int, 0)
	if root == nil {
		return leavel_result
	}
	level_helper(root, 0)
	return leavel_result
}

func level_helper(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if len(leavel_result) == level {
		leavel_result = append(leavel_result, []int{})
	}
	leavel_result[level] = append(leavel_result[level], root.Val)
	if root.Left != nil {
		level_helper(root.Left, level+1)
	}
	if root.Right != nil {
		level_helper(root.Right, level+1)
	}
}

func levelOrder_2(root *TreeNode) [][]int {
	level_result := [][]int{}
	q := []*TreeNode{root}
	for len(q) != 0 {
		len_size := len(q)
		tem := []int{}
		for i := 0; i < len_size; i++ {
			temp := q[0]
			q = q[1:]
			tem = append(tem, temp.Val)
			if temp.Left != nil {
				q = append(q, temp.Left)
			}
			if temp.Right != nil {
				q = append(q, temp.Right)
			}
		}
		level_result = append(level_result, tem)
	}
	return level_result
}
