package leadcode

func numSquares(n int) int {
	dp:=make([]int,n+1)
	for i:=1;i<=n;i++{
		dp[i] = i
		for j:=1;i-j*j>=0;j++{
			dp[i] = min_sqyare(dp[i],dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min_sqyare(x,y int) int{
	if  x<y{
		return x
	}
	return y
}