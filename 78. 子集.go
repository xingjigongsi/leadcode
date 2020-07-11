package leadcode

/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/

var subset_result [][]int
var k int

func subsets(nums []int) [][]int {
	subset_result = make([][]int, 0)
	item := []int{}
	for k = 0; k < len(nums)+1; k++ {
		subsets_helper(nums, 0, item)
	}
	return subset_result
}

func subsets_helper(nums []int, start int, item []int) {
	if k == len(item) {
		itm := make([]int, len(item))
		copy(itm, item)
		subset_result = append(subset_result, itm)
		return
	}
	for i := start; i < len(nums); i++ {
		item = append(item, nums[i])
		subsets_helper(nums, i+1, item)
		item = item[:len(item)-1]
	}
}

// 迭代法
func subsets_1(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{})
	for _, v := range nums {
		temp := make([][]int, len(result))
		for key, v1 := range result {
			v1 = append(v1, v)
			temp[key] = append(temp[key], v1...)
		}
		for _, v := range temp {
			result = append(result, v)
		}
	}
	return result
}

func Test_subsets(nums []int) [][]int {
	return subsets_1(nums)
}
