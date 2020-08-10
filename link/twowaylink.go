package link

import (
	"fmt"
	"strconv"
)

type Node struct {
	key int
	value int
	pre *Node
	next *Node
}

type DoubleList struct {
	capsize int
	size int
	head *Node
	tail *Node
}


func CreateNode(key,value int) *Node{
	return &Node{key:key,value:value}
}

func CreateDoubleList(capsize int) *DoubleList{
	return &DoubleList{
		capsize:capsize,
	}
}

// 从头部添加节点
func (this *DoubleList) Addfirst(node *Node) *Node{
	if this.size>this.capsize{
		return nil
	}
	if this.head == nil{
		this.head = node
		this.tail = node
		this.head.pre = nil
		this.head.next = nil
	}else{
		node.next = this.head
		this.head.pre = node
		this.head = node
		this.head.pre = nil
	}
	this.size++
	return node
}

func (this *DoubleList) AddLast(node *Node) *Node{
	if this.size>this.capsize{
		return nil
	}
	if this.tail == nil{
		this.tail = node
		this.head = this.tail
		this.tail.pre = nil
		this.tail.next = nil
	}else{
		this.tail.next = node
		node.pre = this.tail
		this.tail = node
		this.tail.next = nil
	}
	this.size++
	return node
}

// 任意节点删除
func (this *DoubleList) RemoveNode(node *Node){
	if node == nil{
		this.tail = nil
	}else if node == this.head{
		// 删除头部
	}else if node == this.tail{
		// 删除尾部
	}else{
		// 删除任意节点
		node.next.pre = node.pre
		node.pre.next = node.next
		this.size--
	}

}

//删除头部
func (this *DoubleList) RemoveHead() *Node{
	fmt.Println("fdfddf")
	if this.head == nil{
		return nil
	}
	node:= this.head
	if node.next!=nil{
		this.head = node.next
		this.head.pre = nil
	}else{
		this.head = nil
		this.tail = nil
	}
	this.size--
	return node
}

//删除尾部
func (this *DoubleList) RemoveTail() *Node{
	if this.tail == nil{
		return nil
	}
	node := this.tail
	if node.pre!=nil{
		this.tail = node.pre
		this.tail.next = nil
	}else{
		this.head = nil
		this.tail = nil
	}
	this.size--
	return node
}

func (this *DoubleList) String() string{
	line:=""
	p:=this.head
	for p!=nil{
		line += strconv.Itoa(p.key)+"->"
		p = p.next
	}
	if line!=""{
		line+="nil"
	}
	return line
}


