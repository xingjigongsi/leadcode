package leadcode

/**
实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/powx-n
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/

func myPow(x float64, n int) float64 {
	if n > 0 {
		return myPow_helper(x, n)
	}
	return 1.0 / myPow_helper(x, -n)
}

func myPow_1(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	} else {
		x = x / 1
		n = -n
	}
	if n%2 == 0 {
		return myPow_1(x*x, n/2)
	} else {
		return x * myPow_1(x*x, n/2)
	}
}

func myPow_helper(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	re := myPow_helper(x, n/2)
	if n%2 == 0 {
		return re * re
	}
	return re * re * x
}

func myPow_helper_1(x float64, n int) float64 {
	result := 1.0
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n %= 2
	}
	return result
}

func Test_myPow(x float64, n int) float64 {
	return myPow(x, n)
}
