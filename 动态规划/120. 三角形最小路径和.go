package leadcode

import "strconv"

/**
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/

// 自底向上
func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= len(triangle[i])-1; j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func minimumTotal_1(triangle [][]int) int {
	result := []int{}
	result = append(result, triangle[len(triangle)-1]...)
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; i++ {
			result[i] = min(result[i], result[i+1]) + triangle[i][j]
		}
	}
	return result[0]
}

var min_size int
var min_maps map[string]int

func minimumTotal_2(triangle [][]int) int {
	min_size = len(triangle)
	min_maps = make(map[string]int, 0)
	return mini_helper(triangle, 0, 0)
}

func mini_helper(triangle [][]int, level int, c int) int {
	if level == min_size-1 {
		return triangle[level][c]
	}
	key := strconv.Itoa(level) +"_"+strconv.Itoa(c)
	if v, ok := min_maps[key]; ok {
		return v
	}
	left := mini_helper(triangle, level+1, c)
	right := mini_helper(triangle, level+1, c+1)
	ret := min(left, right) + triangle[level][c]
	if _, ok := min_maps[key]; !ok {
		min_maps[key] = ret
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
