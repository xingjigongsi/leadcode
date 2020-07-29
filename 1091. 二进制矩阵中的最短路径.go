package leadcode

import (
	"math"
	"strconv"
)

var short_targets = [][]int{{-1,0},{-1,1},{-1,-1},{0,-1},{1,-1},{1,0},{1,1},{0,1}}
var w,h int
var shortes_result int
var map_shortes map[string]bool

func ShortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0]!=0 || grid[len(grid)-1][len(grid[0])-1]!=0{
		return -1
	}
	w = len(grid)
	h = len(grid[0])
	map_shortes = make(map[string]bool,0)
	shortes_result = math.MaxInt64
	map_shortes[strconv.Itoa(0)+"_"+strconv.Itoa(0)] = true
	shortest_helper(grid,0,0,1)
	if shortes_result==math.MaxInt64{
		return -1
	}
	return shortes_result
}

func shortest_helper(grid [][]int,x int,y int,temp int){
	if x==w-1 && y==h-1{
		shortes_result = min_shortes(shortes_result,temp)
		return
	}
	for _,v:=range short_targets{
		new_x:= x+v[0]
		new_y := y+v[1]
		if checkWidth(new_x,new_y) && grid[new_x][new_y]==0{
			key1:= strconv.Itoa(new_x)+"_"+strconv.Itoa(new_y)
			if _,ok:= map_shortes[key1];!ok{
				map_shortes[key1] = true
				shortest_helper(grid,new_x,new_y,temp+1)
				delete(map_shortes,key1)
			}
		}
	}
}

func checkWidth(x,y int) bool{
	return x>=0 && x<w && y<h && y>=0
}

func min_shortes(x,y int) int{
	if x>y{
		return y
	}
	return x
}

// 没有最优子结构
func shortestPathBinaryMatrix_1(grid [][]int) int {
	if grid[0][0]!=0 || grid[len(grid)-1][len(grid[0])-1]!=0{
		return -1
	}
	w := len(grid)
	h := len(grid[0])
	dp:=make([][]int,w+1)
	for i,_:=range dp{
		dp[i] = make([]int,h+1)
	}
	var short_targets = [][]int{{-1,0},{-1,1},{-1,-1},{0,-1},{1,-1},{1,0},{1,1},{0,1}}
	temp:= math.MaxInt64
	for i:=0;i<w;i++{
		for j:=0;j<h;j++{
			for _,v:=range short_targets{
				x_i:=i+v[0]
				y_j:=j+v[1]
				if checkWidth(x_i,y_j){
					dp[i][j] = min_shortes(temp,dp[i][j]+dp[x_i][y_j])
					temp = dp[i][j]
				}
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}



func shortestPathBinaryMatrix_2(grid [][]int) int {
	if grid[0][0]!=0 || grid[len(grid)-1][len(grid[0])-1]!=0{
		return -1
	}
	w = len(grid)
	h = len(grid[0])
	vistied:= make([][]bool,w+1)
	for i,_:=range vistied{
		vistied[i] = make([]bool,h+1)
	}
	vistied[0][0] = true
	q:=[][]int{{0,0}}
	var short_targets = [][]int{{-1,0},{-1,1},{-1,-1},{0,-1},{1,-1},{1,0},{1,1},{0,1}}
	dep:= 0
	for len(q)>0{
		len_q:= len(q)
		for i:=0;i<len_q;i++{
			temp:= q[0]
			if temp[0]== w-1 && temp[1]== h-1{
				return dep+1
			}
			for _,v:=range short_targets{
				new_x:= temp[0]+v[0]
				new_y := temp[1]+v[1]
				if checkWidth(new_x,new_y) && grid[new_x][new_y]==0 && !vistied[new_x][new_y]{
					vistied[new_x][new_y] = true
					q = append(q,[]int{new_x,new_y})
				}

			}
			q =q[1:]
		}
		dep++
	}
	return -1
}

