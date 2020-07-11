package leadcode

// 第一种
func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	key := 0
	min_height := 0
	for {
		if i >= j {
			break
		}
		key = j - i
		if height[i] <= height[j] {
			min_height = height[i]
			i++

		} else if height[i] > height[j] {
			min_height = height[j]
			j--

		}
		area := key * min_height
		if max < area {
			max = area
		}
	}
	return max
}

func maxArea1(height []int) int {
	var max, i, j, min_height = 0, 0, len(height) - 1, 0
	for {
		if i >= j {
			break
		}
		if height[i] >= height[j] {
			min_height = height[j]
		} else {
			min_height = height[i]
		}
		area := (j - i) * min_height
		if max < area {
			max = area
		}
		if height[i] >= height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}
