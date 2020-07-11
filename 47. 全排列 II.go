package leadcode

import (
	"sort"

)

/**
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
**/
var peru_result [][]int
var msps map[int]bool

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	peru_result = make([][]int, 0)
	item := make([]int, 0)
	msps = make(map[int]bool, 0)
	permute_helpers(nums, item)
	return peru_result

}

func permute_helpers(nums []int, item []int) {
	if len(nums) == len(item) {
		itm := make([]int, len(item))
		copy(itm, item)
		peru_result = append(peru_result, itm)
		return
	}
	for i, v := range nums {
		if _, ok := msps[i]; !ok {
			if _, ok := msps[i-1]; !ok && i > 0 && nums[i] == nums[i-1] {
				continue
			}
			msps[i] = true
			item = append(item, v)
			permute_helpers(nums, item)
			item = item[:len(item)-1]
			delete(msps, i)

		}
	}

}

func Test_permuteUnique(nums []int) [][]int {
	return permuteUnique(nums)
}
