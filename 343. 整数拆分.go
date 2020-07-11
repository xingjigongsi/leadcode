package leadcode

/**
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
示例 1:
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/
var nums map[int]int

func integerBreak(n int) int {
	nums = make(map[int]int, 0)
	return interger_helper(n)
}

func interger_helper(n int) int {
	if n == 1 {
		return 1
	}
	if _, ok := nums[n]; ok {
		return nums[n]
	}
	inte_result := -1
	for i := 1; i <= n-1; i++ {
		inte_result = maxs(inte_result, i*(n-i), i*interger_helper(n-i))
	}
	if _, ok := nums[n]; !ok {
		nums[n] = inte_result
	}
	return inte_result
}

func maxs(x, y, z int) int {
	max := x
	if max < y {
		max = y
	}
	if max < z {
		max = z
	}
	return max
}

func integerBreak_1(n int) int {
	var verctor = make([]int, n+1)
	verctor[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			verctor[i] = maxs(verctor[i], j*(i-j), j*verctor[i-j])
		}
	}
	return verctor[n]
}
