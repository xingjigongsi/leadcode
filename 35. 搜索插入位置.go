package leadcode

func searchInsert(nums []int, target int) int {
	nums_len := len(nums)
	if nums_len == 0{
		return 0
	}
	if nums[nums_len-1]<target{
		return nums_len
	}
	l,r:=0,nums_len-1
	for l<r{
		m:=l+(r-l)/2
		if nums[m]<target{
			l = m+1
		}else{
			r = m
		}
	}
	return l
}
