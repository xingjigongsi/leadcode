package leadcode

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var binary_result []string
func binaryTreePaths(root *TreeNode) []string {
	binary_result = make([]string,0)
	helper_TreePaths(root,"")
	return binary_result
}

func helper_TreePaths(root *TreeNode,str string){
	if root == nil{
		return
	}
	if root.Left==nil && root.Right==nil{
		str+=strconv.Itoa(root.Val)
		binary_result = append(binary_result,str)
		return
	}
	helper_TreePaths(root.Left,str+strconv.Itoa(root.Val)+"->")
	helper_TreePaths(root.Right,str+strconv.Itoa(root.Val)+"->")
}


func binaryTreePaths_1(root *TreeNode) []string{
	var res []string
	if root == nil{
		return res
	}
	fmt.Println("start ",root.Val,res)
	if root.Left==nil && root.Right == nil{
		res = append(res,strconv.Itoa(root.Val))
		fmt.Println("dp fish ",res)
		return res
	}
	le:= binaryTreePaths(root.Left)
	fmt.Println("left ",le)
	for _,v:=range le{
		res = append(res,strconv.Itoa(root.Val)+"->"+v)
	}
	re:= binaryTreePaths(root.Right)
	fmt.Println("right ",re)
	for _,v:=range re{
		res = append(res,strconv.Itoa(root.Val)+"->"+v)
	}
	fmt.Println("complete ",res)
	return res
}
