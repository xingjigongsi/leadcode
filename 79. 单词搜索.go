package leadcode

/**
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
**/
var tar_gets [][]int = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var visted [][]bool
var x int
var y int

func exist(board [][]byte, word string) bool {
	visted = make([][]bool, len(board))
	x = len(board)
	y = len(board[0])
	for i := 0; i < len(visted); i++ {
		visted[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if exist_helper(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func contact(new_x, new_y int) bool {
	if new_x >= 0 && new_x < x && new_y >= 0 && new_y < y {
		return true
	}
	return false
}

func exist_helper(board [][]byte, word string, index int, start_x int, start_y int) bool {
	if index == len(word)-1 {
		return board[start_x][start_y] == word[index]
	}
	// 四个方向
	if board[start_x][start_y] == word[index] {
		visted[start_x][start_y] = true
		for i := 0; i < len(tar_gets); i++ {
			new_x := start_x + tar_gets[i][0]
			new_y := start_y + tar_gets[i][1]
			if contact(new_x, new_y) && !visted[new_x][new_y] {
				if exist_helper(board, word, index+1, new_x, new_y) {
					return true
				}
			}

		}
		visted[start_x][start_y] = false
	}
	return false
}

func Test_exist(board [][]byte, word string) bool {
	return exist(board, word)
}
