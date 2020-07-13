package leadcode

// K 是股票的买卖次数,是无穷大的
func maxProfit_M(prices []int) int {
	n_Profit:=len(prices)
	if n_Profit==0{
		return 0
	}
	dp:=make([][]int,n_Profit)
	for i,_:=range dp{
		dp[i] = make([]int,2)
	}
	for i:=0;i<n_Profit;i++{
		if i-1<0{
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max_profits(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1] = max_profits(dp[i-1][1],dp[i-1][0]-prices[i])
	}
	return dp[n_Profit-1][0]
}

func max_profits(x,y int) int{
	if x>y{
		return x
	}
	return y
}
