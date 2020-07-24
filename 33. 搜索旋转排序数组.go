package leadcode

func search(nums []int, target int) int {
	nums_len := len(nums)
	if nums_len == 0 {
		return 0
	}
	left,right:=0,nums_len-1
	for left<right{
		m:=left+(right-left)/2
		if nums[0]<=nums[m] && (nums[0]>target || nums[m]<target){
			left = m+1
		}else if nums[0]>target && nums[m]<target{
			left = m+1
		}else{
			right = m
		}
	}
	if left==right && nums[left] == target{
		return left
	}
	return -1
}
