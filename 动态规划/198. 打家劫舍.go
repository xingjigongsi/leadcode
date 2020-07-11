package leadcode

func rob(nums []int) int {
	var n = len(nums)
	if n == 0 {
		return 0
	}
	var maps = make(map[int]int, 0)
	maps[n-1] = n - 1
	temp := 0
	for i := n - 2; i >= 0; i++ {
		for j := i; j < n; j++ {
			if j+2 <= n {
				temp = maps[j+2]
			}
			maps[i] = rob_max(maps[i], maps[j]+temp)
		}
	}
	return maps[0]
}

/**
第二种
**/
func rob_second(nums []int) int {
	var n = len(nums)
	if n == 0 {
		return 0
	}
	var rob_list = make([]int, n)
	rob_list[0] = nums[0]
	rob_list[1] = rob_max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		rob_list[i] = rob_max(nums[i-1], nums[i-2]+nums[i])
	}
	return rob_list[n-1]
}

func rob_max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
