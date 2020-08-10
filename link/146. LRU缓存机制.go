package link

import (
	"strconv"
)


type LRUCache struct {
	doubleList *DoubleList
	nodemap map[int]*Node
}


func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		doubleList:CreateDoubleList(capacity),
		nodemap:make(map[int]*Node,0),
	}
}


func (this *LRUCache) Get(key int) int {
	if v,ok:=this.nodemap[key];ok{
		this.doubleList.RemoveNode(v)
		this.doubleList.Addfirst(v)
		return v.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if v,ok:=this.nodemap[key];ok{
		this.doubleList.RemoveNode(v)
		v.value = value
		this.doubleList.Addfirst(v)
	}else{
		node:= CreateNode(key,value)
		if this.doubleList.size>=this.doubleList.capsize{
			node1:=this.doubleList.RemoveTail()
			delete(this.nodemap,node1.key)
		}
		this.doubleList.Addfirst(node)
		this.nodemap[key] = node
	}
}



//双向链表
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
func (this *DoubleList) RemoveNode(node *Node) *Node{
	if node == nil{
		this.tail = nil
	}else if node == this.head{
		// 删除头部
		return this.RemoveHead()
	}else if node == this.tail{
		// 删除尾部
		return this.RemoveTail()
	}else{
		// 删除任意节点
		node.next.pre = node.pre
		node.pre.next = node.next
		this.size--
		return node
	}
	return node
}

//删除头部
func (this *DoubleList) RemoveHead() *Node{
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


func (this *LRUCache) String() string{
	line:=""
	p:=this.doubleList.head
	for p!=nil{
		line += strconv.Itoa(p.key)+"->"
		p = p.next
	}
	if line!=""{
		line+="nil"
	}
	return line
}