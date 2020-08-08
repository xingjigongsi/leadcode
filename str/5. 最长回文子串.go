package str

// 动态规划 二维数组
func longestPalindrome(s string) string {
	len_s := len(s)
	if len_s<2{
		return s
	}
	dp:= make([][]bool,len_s)
	for i:=0;i<len_s;i++{
		dp[i] = make([]bool,len_s)
		dp[i][i] = true
	}
	max_len:=1
	start:=0
	for j:=0;j<len_s;j++{
		for i:=0;j<i;j++{
			if s[i]!=s[j]{
				dp[i][j] = false
			}else{
				if j-i<3{   // 边界情况   j-1-(i+1)+1>2 =>j-i>3
					dp[i][j] = true
				}else{
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 开始计算
			if dp[i][j] && j-i+1>max_len{
				max_len = j-i+1
				start = i
			}
		}
	}
	re:= s[start:max_len]
	return re
}
