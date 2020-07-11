package leadcode

import "math"

// dp
func coinChange(coins []int, amount int) int {
	nums := amount + 1
	var dp []int = make([]int, amount+1)
	for i := 1; i <= nums; i++ {
		dp[i] = amount
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			dp[i] = coin_min(dp[i], dp[i-coins[i]]) + 1
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func coin_min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
方法二  递归
**/
var coin_money map[int]int

func coinChange_1(coins []int, amount int) int {
	coin_money = make(map[int]int, 0)
	return coin_helper(coins, amount)
}

func coin_helper(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if _, ok := coin_money[amount]; ok {
		return coin_money[amount]
	}
	max_value := math.MaxInt64
	for i := 0; i < len(coins); i++ {
		t := coin_helper(coins, amount-coins[i])
		if t >= 0 && t < max_value {
			max_value = t + 1
		}
	}
	if _, ok := coin_money[amount]; !ok {
		if max_value == math.MaxInt64 {
			coin_money[amount] = -1
		} else {
			coin_money[amount] = max_value
		}

	}
	return coin_money[amount]
}
