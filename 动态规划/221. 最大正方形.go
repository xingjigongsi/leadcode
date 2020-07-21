package leadcode

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0{
		return 0
	}
	w_len:=len(matrix)
	h_len:= len(matrix[0])
	dp:=make([][]int,w_len+1)
	for i,_:= range dp{
		dp[i] = make([]int,h_len+1)
	}
	max_width :=0
	for i:=0;i<w_len;i++{
		for j:=0;j<h_len;j++{
			if matrix[i][j] == '1'{
				dp[i+1][j+1] = Square_min(dp[i][j+1],dp[i][j],dp[i+1][j])+1
				max_width = squere_max(max_width,dp[i+1][j+1])
			}
		}
	}
	return max_width*max_width
}

func Square_min(x,y,z int) int{
	min:=x
	if min>y{
		min = y
	}
	if min>z{
		min = z
	}
	return min
}

func squere_max(x,y int) int{
	if x>y{
		return x
	}
	return y
}

