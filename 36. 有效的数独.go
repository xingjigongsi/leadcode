package leadcode

import "strconv"

func isValidSudoku(board [][]byte) bool {
	rows:= make(map[string]int,9)
	cols:= make(map[string]int,9)
	boxes:= make(map[string]int,9)
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if board[i][j]=='.'{
				continue
			}
			key1:= strconv.Itoa(i)+"_"+string(board[i][j])
			rows[key1]+=1
			key2:=strconv.Itoa(j)+"_"+string(board[i][j])
			cols[key2]+=1
			keys3:=strconv.Itoa((i/3)*3+j/3)+string(board[i][j])
			boxes[keys3]+=1
			if rows[key1]>1 || cols[key2]>1 || boxes[keys3]>1{
				return false
			}
		}
	}
	return true
}



func isValidSudoku_1(board [][]byte) bool {
	var rows,cols,boxes [9][9]int
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if board[i][j]=='.'{
				continue
			}
			nums:= board[i][j]-'1'
			rows[i][nums]+=1
			cols[j][nums]+=1
			boxes[(i/3)*3+j/3][nums]+=1
			if rows[i][nums]>1 || cols[j][nums]>1 || boxes[(i/3)*3+j/3][nums]>1{
				return false
			}
		}
	}
	return true
}



func isValidSudoku_4(board [][]byte) bool {
	var row, col, block [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '1')  // 当前数字的 二进制数位 位置
			bi := i/3 + j/3*3  // 3x3的块索引号
			if (row[i] & cur) | (col[j] & cur) | (block[bi] & cur) != 0 { // 使用与运算查重
				return false
			}
			// 在对应的位图上，加上当前数字
			row[i] |= cur
			col[j] |= cur
			block[bi] |= cur
		}
	}
	return true
}