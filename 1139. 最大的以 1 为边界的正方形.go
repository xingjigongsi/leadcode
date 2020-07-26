package leadcode

//

func largest1BorderedSquare(grid [][]int) int {
	if len(grid) == 0 || len(grid[0])==0 {
		return 0
	}
	grid_w:=len(grid)
	grid_h:=len(grid[0])
	dp:=make([][]int,grid_w+1)
	for i,_:= range dp{
		dp[i] = make([]int,grid_h+1)
	}
	max_width:=0
	for i:=0;i<grid_w;i++{
		for j:=0;j<grid_h;j++{
			if grid[i][j] ==1{
				dp[i+1][j+1] = large_max(dp[i+1][j],dp[i][j+1],dp[i-1][j-1])+1
				max_width = max_d(dp[i+1][j+1],max_width)
			}
			if i-1<0 || j-1<0{
				continue
			}
			if grid[i][j]==0 && grid[i-1][j]==1 && grid[i][j-1] == 1 && grid[i-1][j-1]==1{
				dp[i+1][j+1] = large_max(dp[i+1][j],dp[i][j+1],dp[i-1][j-1])+1
				max_width = max_d(dp[i+1][j+1],max_width)
			}
		}
	}
	return max_width*max_width
}

func large_max(x,y,z int) int{
	max:=x
	if max<y{
		max = y
	}
	if max<z{
		max = z
	}
	return max
}

func max_d(x,y int) int{
	if x<y{
		return y
	}
	return x
}