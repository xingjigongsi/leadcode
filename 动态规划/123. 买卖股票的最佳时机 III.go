package leadcode

func maxProfit_N(prices []int) int {
	n_profit:=len(prices)
	if n_profit==0{
		return 0
	}
	dp:=make([][][]int,n_profit)
	for i,_:=range dp{
		dp[i] = make([][]int,2)

	}
	for i:=0;i<n_profit;i++{

	}
	return dp[n_profit-1][0]
}


func m_max(x,y int) int{
	if x>y{
		return x
	}
	return y
}
