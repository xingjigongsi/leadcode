package leadcode

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sum_map:=make(map[int]bool)
	for i:=0;i<len(A);i++{
		for j:=0;j<len(B);j++{
			sum_map[A[i]+B[i]] = true
		}
	}
	result:=0
	for i:=0;i<len(C);i++{
		for j:=0;j<len(D);j++{
			if _,ok:=sum_map[0-(C[i]+D[j])];ok{
				result++
			}
		}
	}
	return result
}
