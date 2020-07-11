package leadcode

import (
	"strings"

)

/**
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
**/
var solve_result [][]string
var cols []bool
var left_cols []bool  // 对角线的坐标
var right_cols []bool // 对角线的坐标

func solveNQueens(n int) [][]string {
	solve_result = make([][]string, 0)
	item := []int{}
	cols = make([]bool, n)
	left_cols = make([]bool, 2*n-1)
	right_cols = make([]bool, 2*n-1)

	solves_helper(n, 0, item)
	return solve_result
}

func solves_helper(n int, index int, item []int) {
	if n == index {
		itm := make([]int, len(item))
		copy(itm, item)
		solve_result = append(solve_result, changestring(n, itm))
		return
	}
	for i := 0; i < n; i++ {
		if !cols[i] && !left_cols[index+i] && !right_cols[index-i+n-1] {
			cols[i] = true
			left_cols[index+i] = true
			right_cols[index-i+n-1] = true
			item = append(item, i)
			solves_helper(n, index+1, item)
			cols[i] = false
			left_cols[index+i] = false
			right_cols[index-i+n-1] = false
			item = item[:len(item)-1]
		}

	}
}

func changestring(n int, item []int) []string {
	res := []string{}
	str := strings.Repeat(".", n)
	for i := 0; i < n; i++ {
		res = append(res, str)
	}
	for i := 0; i < len(item); i++ {
		c := []byte(res[i])
		c[item[i]] = 'Q'
		res[i] = string(c)
	}
	return res
}

func Test_solveNQueens(n int) [][]string {
	return solveNQueens(n)
}
