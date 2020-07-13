package leadcode

import "math"

func maxProfit_NI(k int, prices []int) int {
	profit_n := len(prices)
	if profit_n==0{
		return 0
	}
	if k>profit_n/2{
		return max_pro_fit(prices)
	}
	dp:=make([][][]int,profit_n)
	for i,_:=range dp{
		dp[i] = make([][]int,2)
		dp[i][0] = make([]int,k+1)
		dp[i][1] = make([]int,k+1)
	}
	for i:=0; i<profit_n; i++ {
		for j:=k;j>=1;j--{
			if i-1==-1{
				dp[i][0][j] = 0
				dp[i][1][j] = -prices[i]
				continue
			}
			dp[i][0][j] = max_pforit(dp[i-1][0][j],dp[i-1][1][j]+prices[i])
			dp[i][1][j] = max_pforit(dp[i-1][1][j],dp[i-1][0][j-1]-prices[i])
		}
	}
	return dp[profit_n-1][0][k]
}

func max_pro_fit(prices []int) int{
	dp_i_0,dp_i_1:=0,math.MinInt16
	for i:=0;i<len(prices);i++{
		dp_i_0 = max_pforit(dp_i_0,dp_i_1+prices[i])
		dp_i_1 = max_pforit(dp_i_1,dp_i_0-prices[i])
	}
	return dp_i_0
}

func max_pforit(x,y int) int{
	if x>y{
		return x
	}
	return y
}