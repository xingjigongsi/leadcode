package leadcode


func solveSudoku(board [][]byte)  {
	if len(board) == 0{
		return
	}
	solve(board)
}

func solve(board [][]byte) bool  {
	var c byte
	for i,_:=range board{
		for j,v:= range board[i]{
			if v != '.'{
				continue
			}
			for c='1';c<='9';c++{
				if isSove(board,i,j,c){
					board[i][j] = c
					if solve(board){
						return true
					}else{
						board[i][j] = '.'
					}
				}
			}
			return false
		}
	}
	return true
}

// 判斷 数组是否合法
func  isSove(board [][]byte,row,col int,c byte) bool{
	for i:=0;i<9;i++{
		if board[i][col]!='.' && board[i][col] == c {
			return false
		}
		if board[row][i]!='.' && board[row][i] == c{
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3]!='.' && board[3*(row/3)+i/3][3*(col/3)+i%3] == c{
			return false
		}
	}
	return true
}
