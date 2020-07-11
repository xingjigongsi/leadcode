package leadcode

/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

示例 1:

输入:
11110
11010
11000
00000
输出: 1
示例 2:

输入:
11000
11000
00100
00011
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/
var m int
var n int
var num_vistied [][]bool
var targets [][]int = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func numIslands(grid [][]byte) int {
	var result int
	m = len(grid)
	if m == 0 {
		return 0
	}
	n = len(grid[0])
	num_vistied = make([][]bool, m)
	for i := 0; i < m; i++ {
		num_vistied[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !num_vistied[i][j] {
				result++
				num_helper(grid, i, j)
			}
		}
	}
	return result
}

func num_helper(grid [][]byte, x, y int) {
	num_vistied[x][y] = true
	for i := 0; i < len(targets); i++ {
		new_x := x + targets[i][0]
		new_y := y + targets[i][1]
		if show_page(new_x, new_y) && grid[new_x][new_y] == '1' && !num_vistied[new_x][new_y] {
			num_helper(grid, new_x, new_y)
		}
	}
}

func show_page(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func Test_numIslands(grid [][]byte) int {
	return numIslands(grid)
}
