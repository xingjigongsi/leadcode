package uniod

import "fmt"

type unio_1 struct {
	data []int
}

func InitUnio_1(n int) *unio_1{
	temp:= unio_1{
		data:make([]int,n),
	}
	for i:=0;i<n;i++{
		temp.data[i] = i
	}
	return &temp
}

func (this *unio_1) findPrent(p int) int{

	for p!=this.data[p]{
		p = this.data[p]
	}
	return this.data[p]
}

// 是否链接
func (this *unio_1) IsConnect(p,q int) bool{
	pPrent:=this.findPrent(p)
	qPrent:=this.findPrent(q)
	return pPrent == qPrent
}

// 合并
func (this *unio_1) UnConnected(p,q int){
	p_prent:=this.findPrent(p)
	q_prent:=this.findPrent(q)
	this.data[p_prent] = q_prent
}

// 显示并查集
func (this *unio_1) Show_uniod1() string{
	str:=""
	for k,v:=range this.data{
		str+=fmt.Sprintf("%d:%d,",k,v)
	}
	return str
}