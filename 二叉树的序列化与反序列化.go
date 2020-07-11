/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leadcode

import (
	"strconv"
	"strings"

)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	statck := []*TreeNode{root}
	result := ""
	for len(statck) != 0 {
		temp := statck[0]
		statck = statck[1:]
		if temp == nil {
			result += "null,"
			continue
		}
		result += strconv.Itoa(temp.Val) + ","
		statck = append(statck, temp.Left, temp.Right)
	}
	return strings.Trim(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	re := strings.Split(data, ",")
	val_temp, _ := strconv.Atoi(re[0])
	root := &TreeNode{Val: val_temp}
	static := []*TreeNode{root}
	for i := 1; i < len(re); i++ {
		if len(static) == 0 {
			break
		}
		parent := static[0]
		static = static[1:]
		if re[i] != "null" {
			val_temp, _ = strconv.Atoi(re[i])
			left := &TreeNode{Val: val_temp}
			static = append(static, left)
			parent.Left = left
		}
		i = i + 1
		if re[i] != "null" {
			val_temp, _ = strconv.Atoi(re[i])
			right := &TreeNode{Val: val_temp}
			static = append(static, right)
			parent.Right = right
		}

	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
