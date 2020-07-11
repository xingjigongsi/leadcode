package leadcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	len_nums := len(nums) - 1
	result := [][]int{}
	for i := 0; i <= len_nums; i++ {
		l, r := i+1, len_nums
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len_nums && nums[r] == nums[r+1] {
				r--
				continue
			}
			if nums[l]+nums[r]+nums[i] == 0 {
				temp := []int{nums[i], nums[l], nums[r]}
				result = append(result, temp)
				l++
				r--
			} else if nums[l]+nums[r]+nums[i] > 0 {
				r--
			} else {
				l++
			}

		}
	}
	return result
}

//三层循坏
func threeSum1(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if i > 0 && nums[i] == nums[i-1] {
					continue
				}
				if j > i+1 && nums[j] == nums[j-1] {
					continue
				}
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					result = append(result, temp)
				}
			}
		}
	}
	return result
}

// hash 表 (error)
func threeSum2(nums []int) [][]int {
	result := [][]int{}
	res := make(map[int][]int)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			res[nums[i]-nums[j]] = []int{nums[i], nums[j]}
		}
	}
	for k := 0; k < len(nums)-1; k++ {
		kk, ok := res[nums[k]]
		if ok && kk[0]+kk[1]+nums[k] == 0 {
			temp := []int{kk[0], kk[1], nums[k]}
			result = append(result, temp)
		}
	}
	return result
}
