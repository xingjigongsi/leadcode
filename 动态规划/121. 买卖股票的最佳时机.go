package leadcode

// 买，卖，持有 0 是卖出   1 是 持有
func maxProfit(prices []int) int {
	n_prices := len(prices)
	if n_prices==0{
		return 0
	}
	dp:=make([][][]int,n_prices)
	for i,_:=range dp{
		dp[i] = make([][]int,2)
		dp[i][0] = make([]int,3)
		dp[i][1] = make([]int,3)
	}
	for i:=0;i<len(prices);i++{
		for j:=1;j<=3;j++{
			if i-1<0{
				dp[0][0][j] = 0
				dp[0][1][j] = -prices[i]
				continue
			}
			dp[i][0][j] = pro_max(dp[i-1][0][j],dp[i-1][1][j]+prices[i])
			dp[i][1][j] = pro_max(dp[i-1][1][j],dp[i-1][0][j-1]-prices[i])
		}

	}
	return dp[n_prices-1][0][3]
}

// 第二种二维数组 (0 没有交易  1卖出 2 买入 3卖出 4 买入)
func maxProfit_III(prices []int) int {
	n_prices := len(prices)
	if n_prices==0{
		return 0
	}
	dp:=make([][]int,n_prices)
	for i,_:=range dp{
		dp[i] = make([]int,5)
	}
	for i:=0;i<n_prices;i++{
		if i-1==-1{
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			dp[i][2] = 0
			dp[i][3] = -prices[i]
			dp[i][4] = 0
			continue
		}
		dp[i][0] = 0
		dp[i][1] = pro_max(dp[i-1][1],dp[i-1][0]-prices[i])
		dp[i][2] = pro_max(dp[i-1][2],dp[i-1][1]+prices[i])
		dp[i][3] = pro_max(dp[i-1][3],dp[i-1][2]-prices[i])
		dp[i][4] = pro_max(dp[i-1][4],dp[i-1][3]+prices[i])
	}
	return pro_max(dp[n_prices-1][2],dp[n_prices-1][4])
}


/*
  第三种一维数组 (0 没有交易  1卖出 2 买入 3卖出 4 买入)
 */
func maxProfit_IIII(prices []int) int {
	n_prices := len(prices)
	if n_prices==0{
		return 0
	}
	dp:=make([]int,5)
	for i:=0;i<n_prices;i++{
		if i-1 == -1{
			dp[0] = 0
			dp[1] = -prices[i]
			dp[2] = 0
			dp[3] = -prices[i]
			dp[4] = 0
			continue
		}
		dp[0] = 0
		dp[1] = pro_max(dp[1],dp[0]-prices[i])
		dp[2] = pro_max(dp[2],dp[1]+prices[i])
		dp[3] = pro_max(dp[3],dp[2]-prices[i])
		dp[4] = pro_max(dp[4],dp[3]+prices[i])
	}
	return pro_max(dp[2],dp[4])
}




func pro_max(x,y int) int{
	if x>y{
		return x
	}
	return y
}