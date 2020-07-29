package uniod

import "fmt"

/**
  路径压缩
 */
type unio_5 struct {
	data []int
	rank []int  // 排名
}

func InitUnio_5(n int) *unio_5{
	temp:= unio_5{
		data:make([]int,n),
		rank:make([]int,n),
	}
	for i:=0;i<n;i++{
		temp.data[i] = i
		temp.rank[i] = 1
	}
	return &temp
}

func (this *unio_5) findPrent(p int) int{

	if p!=this.data[p]{
		this.data[p] = this.findPrent(this.data[p])
	}
	return this.data[p]
}

// 是否链接
func (this *unio_5) IsConnect(p,q int) bool{
	pPrent:=this.findPrent(p)
	qPrent:=this.findPrent(q)
	return pPrent == qPrent
}

// 合并
func (this *unio_5) UnConnected(p,q int){
	p_prent:=this.findPrent(p)
	q_prent:=this.findPrent(q)
	if this.rank[p_prent]<this.rank[q_prent]{
		this.data[p_prent] = q_prent
	}else if this.rank[p_prent]>this.rank[q_prent]{
		this.data[q_prent] = this.data[p_prent]
	}else{
		this.data[q_prent] = p_prent
		this.rank[p_prent]+=1
	}
}

// 显示并查集
func (this *unio_5) Show_uniod1() (string,string){
	str,str1:="",""
	for k,v:=range this.data{
		str+=fmt.Sprintf("%d:%d,",k,v)
	}
	for k,v:=range this.rank{
		str1+=fmt.Sprintf("%d:%d,",k,v)
	}
	return str,str1
}
