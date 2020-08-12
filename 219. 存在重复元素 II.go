package leadcode


// 滑动窗口
func containsNearbyDuplicate(nums []int, k int) bool {
	contains:=make(map[int]int,0)
	for i:=0;i<len(nums);i++{
		if _,ok:=contains[nums[i]];ok{
			return true
		}
		contains[nums[i]] = i
		if len(contains)==k+1{
			delete(contains,nums[i-k])
		}
	}
	return false
}