package leadcode

import "fmt"

/**
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
**/
var usermap map[int]bool
var result [][]int

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	usermap = make(map[int]bool)
	result = make([][]int, 0)
	temp := []int{}
	permute_helper(nums, 0, temp)
	return result
}

func permute_helper(nums []int, index int, temp []int) {
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, temp)
		result = append(result, tmp)
		return 
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(temp)
		if _, ok := usermap[i]; !ok {
			usermap[i] = true
			temp = append(temp, nums[i])
			permute_helper(nums, index+1, temp)
			temp = temp[:len(temp)-1]
			delete(usermap, i)
		}

	}
}

func Test_permute(nums []int) [][]int {
	return permute(nums)
}
