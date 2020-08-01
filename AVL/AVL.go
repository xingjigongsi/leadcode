package AVL

import (
	"math"
)

type avl struct {
	key int
	value string
	height int  //高度
	left *avl
	right *avl
}

func InitAvl(key int,value string) *avl{
	return &avl{
		key:key,
		value:value,
		height:1,
	}
}

/**
* 右旋转      y               x
×           /  \			/ \
×			x  t1   =>      z   y
*                              / \
*          / \                t3  t1
*		  z  t3
*/
func (this *avl) rightRoute(node *avl) *avl{
	x:=node.left
	t3:=x.right
	x.right = node
	node.left = t3
	node.height = max(this.getNodeHeight(node.left),this.getNodeHeight(node.right))+1
	x.height = max(this.getNodeHeight(x.left),this.getNodeHeight(x.right))+1
	return  x
}

/**
*  左旋转
 */
 func (this *avl) leftRoute(node *avl) *avl{
 	x:=node.right
 	t3:=x.left
 	x.left  = node
 	node.right = t3
 	node.height = max(this.getNodeHeight(node.left),this.getNodeHeight(node.right))+1
 	x.height = max(this.getNodeHeight(x.left),this.getNodeHeight(x.right))+1
	 return x
 }

 // 增加节点
func (this *avl) Add(key int,value string) *avl{
	if this==nil{
		return &avl{
			key:key,
			value:value,
			height:1,
		}
	}
	if this.key>key{
		this.left =  this.left.Add(key,value)
	}else if this.key<key{
		this.right =  this.right.Add(key,value)
	}else{
		this.value = value
	}
	//更新节点高度
	this.height = max(this.getNodeHeight(this.left),this.getNodeHeight(this.right))+1
	if this.getHeightDifference(this)>1 && this.left!=nil && this.getHeightDifference(this.left)>=0{
		return this.rightRoute(this)
	}
	if this.getHeightDifference(this)>1 && this.right!=nil && this.getHeightDifference(this.right)>=0{
		return this.leftRoute(this)
	}
	return this
}

//判断是否是一个平衡二叉树
func (this *avl) IsBlance() bool{
	if this == nil{
		return true
	}
	if this.getHeightDifference(this)>1{
		return false
	}
	return this.left.IsBlance() && this.right.IsBlance()
}

//判断是否是二叉搜索树
func (this *avl) IsBst() bool{
	return this.helper(math.MinInt64,math.MaxInt64)
}

//帮助函数
func (this *avl) helper(min int,max int) bool{
	if this == nil{
		return true
	}
	val:= this.key
	if val>= max || val<=min{
		return false
	}
	return this.left.helper(min,val) && this.right.helper(val,max)
}


// 获取节点的高度，包括处理节点为0
func (this *avl) getNodeHeight(node *avl) int{
	if node == nil{
		return 0
	}
	return node.height
}

//左右高度差
func (this *avl) getHeightDifference(node *avl) int {
	return abs(this.getNodeHeight(node.left) - this.getNodeHeight(node.right))
}


func (this *avl) Getkey() int{
	return this.key
}

func (this *avl) GetLeft() *avl{
	return this.left
}

func (this *avl) GetRight() *avl{
	return this.right
}

// 最大值
func max(x,y int) int{
	if x>y{
		return x
	}
	return y
}

// 绝对值
func abs(x int) int{
	if x<0{
		return -x
	}
	return x
}