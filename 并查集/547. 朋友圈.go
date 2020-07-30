package uniod

var parent []int
var rank []int //优先级
func findCircleNum(M [][]int) int {
	if len(M)==0{
		return 0
	}
	parent = make([]int,len(M))
	for i,_:= range parent{
		parent[i] = -1
		rank[i] = 1
	}
	for i:=0;i<len(M);i++{
		for j:=0;j<len(M[0]);j++{
			if M[i][j]==-1 && i!=j{
				union(i,j)
			}
		}
	}
	result:=0
	for _,v:=range parent{
		if v == 1{
			result++
		}
	}
	return result
}

func union(i,j int){
	p_i:= findPrent(i)
	p_j := findPrent(j)
	if rank[i]>rank[j]{
		parent[j] = p_i
	}else if rank[i]<rank[j]{
		parent[i] = p_j
	}else{
		parent[p_j] = i
		rank[i]+=rank[j]
	}
}

func findPrent(i int) int{
	if parent[i] ==-1{
		return i
	}
	if i!=parent[i]{
		parent[i] = findPrent(parent[i])
	}
	return parent[i]
}