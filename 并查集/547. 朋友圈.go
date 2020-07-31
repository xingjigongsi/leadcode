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
	if p_i!=p_j{
		parent[p_i] = parent[p_j]
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


/**
 dfs
 */

var circle_visted []bool
func findCircleNum_one(M [][]int) int {
	if len(M)==0{
		return 0
	}
	result:=0
	circle_visted = make([]bool,len(M))
	for i:=0;i<len(M);i++{
		if !circle_visted[i]{
			result++
			find_dfs(M,i)
		}
	}
	return result
}

func find_dfs(M [][]int,i int){

	for j:=0;j<len(M);j++{
		if M[i][j]==1{
			circle_visted[j] = true
			find_dfs(M,j)
		}
	}
}
