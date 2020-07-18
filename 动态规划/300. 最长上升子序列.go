package leadcode


func lengthOfLIS(nums []int) int {
	len_nums:=len(nums)
	if len_nums ==0{
		return  0
	}
	dp := make([]int,len_nums)
	for i,_:= range dp{
		dp[i] = 1
	}
	for i:=1;i<len_nums;i++{
		for j:=0;j<i;j++{
			if nums[i]>nums[j]{
				dp[i] = length_max(dp[i],dp[j]+1)
			}
		}
	}
	res:=1
	for i:=0;i<len_nums;i++{
		res = length_max(res,dp[i])
	}
	return res
}

func length_max(x,y int) int{
	if x>y{
		return x
	}
	return y
}