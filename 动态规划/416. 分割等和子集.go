package leadcode

import "strconv"

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
var num_map map[string]bool
func canPartition_1(nums []int) bool {
	sum:=0
	num_map = make(map[string]bool,0)
	for _,v:=range nums{
		sum+=v
	}
	if sum%2!=0{
		return false
	}
	C:=sum/2
	return canPar_helper(nums,C,len(nums)-1)
}

func canPar_helper(nums []int,C int,index int) bool{
	if C==0{
		return true
	}
	if C<0 || index<0{
		return false
	}
	key:= strconv.Itoa(C)+"_"+strconv.Itoa(index)
	if v,ok:=num_map[key];ok{
		return v
	}
	t:= canPar_helper(nums,C,index-1) || canPar_helper(nums,C-nums[index],index-1)
	if _,ok:=num_map[key];!ok{
		num_map[key] = t
	}
	return t
}

