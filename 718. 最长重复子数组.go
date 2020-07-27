package leadcode

func findLength(A []int, B []int) int {
	len_a := len(A)
	len_b := len(B)
	if len_a==0 || len_b==0{
		return 0
	}
	dp:=make([][]int,len_a+1)
	for i,_:=range dp{
		dp[i] = make([]int,len_b+1)
	}
	max:=0
	for i:=0;i<len_a;i++{
		for j:=0;j<len_b;j++{
			temp:=0
			if i-1>=0 && j-1>=0{
				temp = dp[i-1][j-1]
			}
			if A[i] == B[j] {
				dp[i][j] = temp+1
			}
			if dp[i][j]>max{
				max = dp[i][j]
			}
		}
	}
	return max
}
