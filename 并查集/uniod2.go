package uniod

import "fmt"

type unio_2 struct {
	data []int
	size []int
}

func InitUnio_2(n int) *unio_2{
	temp:= unio_2{
		data:make([]int,n),
		size:make([]int,n),
	}
	for i:=0;i<n;i++{
		temp.data[i] = i
		temp.size[i] = 1
	}
	return &temp
}

func (this *unio_2) findPrent(p int) int{

	for p!=this.data[p]{
		p = this.data[p]
	}
	return this.data[p]
}

// 是否链接
func (this *unio_2) IsConnect(p,q int) bool{
	pPrent:=this.findPrent(p)
	qPrent:=this.findPrent(q)
	return pPrent == qPrent
}

// 合并
func (this *unio_2) UnConnected(p,q int){
	p_prent:=this.findPrent(p)
	q_prent:=this.findPrent(q)
	if this.size[p_prent]<=this.size[q_prent]{
		this.data[p_prent] = q_prent
		this.size[q_prent]+=this.size[p_prent]
	}else{
		this.data[q_prent] = p_prent
		this.size[q_prent]+=this.size[p_prent]
	}
}

// 显示并查集
func (this *unio_2) Show_uniod1() string{
	str:=""
	for k,v:=range this.data{
		str+=fmt.Sprintf("%d:%d,",k,v)
	}
	return str
}
