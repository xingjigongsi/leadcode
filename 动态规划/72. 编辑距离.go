package leadcode

func minDistance(word1 string, word2 string) int {
	m:=len(word1)
	n:=len(word2)
	dp:=make([][]int,m+1)
	for i,_:=range dp{
		dp[i] = make([]int,n+1)
	}
	for i:=0;i<m+1;i++{
		dp[i][0] = i
	}
	for j:=1;j<n+1;j++{
		dp[0][j] =j
	}
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			if word1[i-1] == word2[j-1]{
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] = Distance_min(dp[i-1][j],dp[i-1][j-1],dp[i][j-1])+1
			}
		}
	}
	return dp[m][n]
}

func Distance_min(x,y,z int) int{
	min:=x
	if min>y{
		min = y
	}
	if min>z{
		min = z
	}
	return min
}