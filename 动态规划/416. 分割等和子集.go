package leadcode


func canPartition(nums []int) bool {
	len_nums:=len(nums)
	sum:=0
	for _,v:=range nums{
		sum+=v
	}
	if sum%2!=0{
		return false
	}
	C:=sum/2
	dp:=make([]bool,C+1)
	for i:=0;i<=C;i++{
		dp[i] = nums[0] == i
	}
	for i:=0;i<len_nums;i++{
		for j:=C;j>=nums[i];j--{
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[C]
}

// 递归
func canPartition_1(nums []int) bool {

}
