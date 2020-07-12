package leadcode

/** 超时傻递归
var result int
var targets = [][]int{{0,1},{1,0}}
var m,n int
func minPathSum(grid [][]int) int {
	result = math.MaxInt16
	m = len(grid)
	n = len(grid[0])
	Path_helper(grid,0,0,grid[0][0])
	return result
}

func Path_helper(grid [][]int,x,y int,res int){
	if x==m-1 && y==n-1{
		result = path_max(result,res)
		return
	}
	m = len(grid)
	n = len(grid[0])
	Path_helper(grid,0,0,grid[0][0])
	return result
	for _,v:=range targets{
		new_x := x+v[0]
		new_y := y+v[1]
		if column(new_x,new_y){
			Path_helper(grid,new_x,new_y,res+grid[new_x][new_y])
		}

	}
}

func column(new_x,new_y int) bool{
	return new_x<m && new_y<n
}

func path_max(x ,y int) int{
	if x>y{
		return y
	}
	return x
}
*/

/* 超时
var m ,n int

func minPathSum(grid [][]int) int {
	m = len(grid)
	n = len(grid[0])
	return Path_helper(grid,0,0)
}

func Path_helper(grid [][]int,x,y int) int{
	if x==m && y==n{
		return math.MaxInt16
	}
	if x==m-1 &&  y==n-1{
		return grid[x][y]
	}
	return grid[x][y]+path_max(Path_helper(grid,x,y+1),Path_helper(grid,x+1,y))
}

func path_max(x ,y int) int{
	if x>y{
		return y
	}
	return x
}
*/

// 第一种  二维dp
func minPathSum(grid [][]int) int {
	m:=len(grid)
	n:=len(grid[0])
	var dp = make([][]int,m)
	for i,_:=range dp{
		dp[i] = make([]int,n)
	}
	for i:=m-1;i>=0;i--{
		for j:=n-1;j>=0;j--{
			if i==m-1 && j!=n-1{
				dp[i][j] = dp[i][j+1]+grid[i][j]
			}else if i!=m-1 && j==n-1{
				dp[i][j] = dp[i+1][j]+grid[i][j]
			}else if i!=m-1 && j!=n-1{
				dp[i][j] = grid[i][j]+path_max(dp[i][j+1],dp[i+1][j])
			}else{
				dp[i][j] = grid[i][j]
			}
		}
	}
	return dp[0][0]
}


// 第一种  一维dp  一行一行的
func minPathSum_1(grid [][]int) int{
	m:=len(grid)
	n:=len(grid[0])
	dp := make([]int,m)
	for i:=m-1;i>=0;i--{
		for j:=n-1;j>=0;j--{
			if i==m-1 && j!=n-1{
				dp[j] = grid[i][j]+dp[j+1]
			}else if i!=m-1 && j==n-1{
				dp[j] = grid[i][j]+dp[j]
			}else if i!=m-1 &&  j!=n-1{
				dp[j] = grid[i][j]+path_max(dp[j+1],dp[j])
			}else{
				dp[j] = grid[i][j]
			}
		}
	}
	return dp[0]
}

// 不需要额外空间
// grid(i,j)=grid(i,j)+min(grid(i+1,j),grid(i,j+1))
func minPathSum_2(grid [][]int) int{
	m:=len(grid)
	n:=len(grid[0])
	for i:=m-1;i>=0;i--{
		for j:=n-1;j>=0;j--{
			if i== m-1 && j!= n-1{
				grid[i][j] = grid[i][j]+grid[i][j+1]
			}else if i!=m-1 && j==n-1{
				grid[i][j] =  grid[i][j]+grid[i+1][j]
			}else if i!=m-1 &&  j!=n-1{
				grid[i][j] = grid[i][j]+path_max(grid[i+1][j],grid[i][j+1])
			}
		}
	}
	return grid[0][0]
}

func path_max(x ,y int) int{
	if x>y{
		return y
	}
	return x
}