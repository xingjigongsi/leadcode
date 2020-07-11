package leadcode

func intersection(nums1 []int, nums2 []int) []int {
	temp := make(map[int]int)
	for _, v := range nums1 {
		temp[v] = v
	}
	reuslt := []int{}
	for _, v := range nums2 {
		if _, ok := temp[v]; ok {
			reuslt = append(reuslt, v)
			delete(temp, v)
		}
	}
	return reuslt
}
