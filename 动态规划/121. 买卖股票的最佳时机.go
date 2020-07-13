package leadcode

import "math"

// 买，卖，持有 0 是卖出   1 是 持有
func maxProfit(prices []int) int {
	n_prices := len(prices)
	dp:=make([][]int,n_prices)
	for i,_:=range dp{
		dp[i] = make([]int,2)
	}
	for i:=0;i<len(prices);i++{
		if i-1<0{
			dp[0][0] = 0
			dp[0][1] = -prices[i]
			continue
		}
		dp[i][0] = pro_max(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1] = pro_max(dp[i-1][1],dp[i-1][0]-prices[i])
	}
	return dp[n_prices-1][0]
}


func maxProfit_1(prices []int) int {
	n_prices := len(prices)
	if n_prices==0{
		return 0
	}
	price_0,price_1 := 0,math.MinInt16
	for i:=0;i<n_prices;i++{
		price_0 = pro_max(price_0,price_1+prices[i])
		price_1 = pro_max(price_1,-prices[i])
	}
	return price_0
}

func pro_max(x,y int) int{
	if x>y{
		return x
	}
	return y
}