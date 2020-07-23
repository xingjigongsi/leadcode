package leadcode


func canJump(nums []int) bool {
	nums_len:= len(nums)
	if nums_len==0{
		return false
	}
	cat_index:=nums_len-1;
	for i:= nums_len-1;i>=0;i--{
		if nums[i]+i>=cat_index{
			cat_index = i
		}
	}
	return cat_index==0
}
