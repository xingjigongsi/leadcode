package leadcode

func intersect(nums1 []int, nums2 []int) []int {
	temp := map[int]int{}
	for _, v := range nums1 {
		if val, ok := temp[v]; !ok {
			temp[v] = 1
		} else {
			temp[v] = val + 1
		}
	}
	result := []int{}
	for _, v := range nums2 {
		if _, ok := temp[v]; ok {
			result = append(result, v)
			temp[v] = temp[v] - 1
			if temp[v] == 0 {
				delete(temp, v)
			}
		}
	}
	return result
}
