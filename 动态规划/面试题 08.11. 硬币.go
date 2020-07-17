package leadcode

func waysToChange(n int) int {
	var money = []int{25,10,5,1}
	dp:=make([]int,n+1)
	dp[1] = 1
	for i:=1;i<=n;i++{
		for _,m:=range money{
			if i-m>=0{
				dp[i] = dp[i]+dp[i-m]
			}else{
				dp[i] = dp[i]+0
			}

		}
	}
	return dp[n]
}
