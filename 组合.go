package leadcode

/**
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
示例:
输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
**/
var com_result [][]int

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return com_result
	}
	com_result = make([][]int, 0)
	temp := []int{}
	com_helper(n, k, 1, temp)
	return com_result
}

func com_helper(n int, k int, start int, temp []int) {
	if k == len(temp) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		com_result = append(com_result, tmp)
		return
	}
	for i := start; i <= n-(k-len(temp))+1; i++ {
		temp := append(temp, i)
		com_helper(n, k, i+1, temp)
		temp = temp[:len(temp)-1]
	}
}
