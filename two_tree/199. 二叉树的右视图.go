package two_tree

var result []int

// 所谓右视图就是最新看见的
func rightSideView(root *TreeNode) []int {
	result = make([]int,0)
	if root ==nil{
		return result
	}
	right_helper(root,0)
	return result
}

func right_helper(root *TreeNode,deep int){
	if root == nil{
		return
	}
	if deep == len(result){
		result = append(result,root.Val)
	}
	deep++
	right_helper(root.Right,deep)
	right_helper(root.Left,deep)
}

func BFSrightSideView(root *TreeNode) []int{
	result:=[]int{}
	if root == nil{
		return result
	}
	statck:=[]*TreeNode{root}
	for len(statck)>0{
		len_statck:=len(statck)
		for i:=0;i<len_statck;i++{
			temp:= statck[0]
			statck = statck[1:]
			if temp.Right!=nil{
				statck = append(statck,temp.Right)
			}
			if temp.Left!=nil{
				statck = append(statck,temp.Left)
			}
			if i==0{
				result = append(result,temp.Val)
			}
		}
	}
	return result
}





