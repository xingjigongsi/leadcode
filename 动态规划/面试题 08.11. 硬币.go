package leadcode

// 全部背包
func waysToChange(n int) int {
	var money = []int{25,10,5,1}
	var dp [][]int = make([][]int,len(money)+1)
	for i,_:=range dp {
		dp[i] = make([]int,n+1)
	}
	for i:=0;i<=len(money);i++{
		dp[i][0] = 1
	}
	for i:=1;i<=len(money);i++{
		for j:=1;j<=n;j++{
			if j-money[i-1]>=0{
				dp[i][j] = dp[i-1][j]+dp[i][j-money[i-1]]
			}else{
				dp[i][j] = dp[i-1][j]
			}

		}
	}
	return dp[len(money)][n]%1000000007
}

func waysToChange_1(n int) int{
	var money = []int{25,10,5,1}
	var dp []int = make([]int,n+1)
	dp[0] = 1
	for i:=1;i<=len(money);i++{
		for j:=1;j<=n;j++{
			if j-money[i-1]>=0{
				dp[j] = dp[j]+dp[j-money[i-1]]
			}
		}
	}
	return dp[n]
}