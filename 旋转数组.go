package leadcode

// 暴力破解法
func rotate1(nums []int, k int) {
	len_nums := len(nums)
	for i := 0; i < k; i++ {
		for j := 0; j < len_nums; j++ {
			nums[j], nums[len_nums-1] = nums[len_nums-1], nums[j]
		}
	}
}

// 使用额外的数组
func rotate2(nums []int, k int) {
	len_nums := len(nums)
	new_nums := make([]int, len_nums)
	for i := 0; i < len_nums; i++ {
		new_nums[(i+k)%len_nums] = nums[i]
	}
	for i := 0; i < len_nums; i++ {
		nums[i] = new_nums[i]
	}
}

// 使用反转
func rotate3(nums []int, k int) {
	len_nums := len(nums) - 1
	// 边界问题
	k %= len(nums)
	reverse(nums, 0, len_nums)
	reverse(nums, 0, k-1)
	reverse(nums, k, len_nums)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
