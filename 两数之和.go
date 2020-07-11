package leadcode

// 第一种  hash 解法
func twoSum1(nums []int, target int) []int {
	result := []int{}
	ma := make(map[int]int)
	for i, v := range nums {
		temp := target - v
		if j, ok := ma[temp]; ok {
			result = append(result, j, i)
		}
		ma[v] = i
	}
	return result
}

// 第二种 暴力破解法
func twoSum2(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); i++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
