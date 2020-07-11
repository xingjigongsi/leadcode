package leadcode

func largestRectangleArea(heights []int) int {
	var len_height = len(heights)
	if len_height == 0 {
		return 0
	}
	if len_height == 1 {
		return heights[0]
	}
	max_area := 0
	stack := []int{}
	for i := 0; i < len_height; i++ {
		for len(stack) != 0 && heights[stack[len(stack)-1]] > heights[i] {
			curheight := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			for len(stack) != 0 && curheight == heights[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}
			curwidth := 0
			if len(stack) == 0 {
				curwidth = i
			} else {
				curwidth = i - stack[len(stack)-1] - 1
			}
			max_area = maxvalue(max_area, curwidth*curheight)
		}
		stack = append(stack, i)
	}

	for len(stack) != 0 {
		var peek_index = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for len(stack) >= 1 && heights[peek_index] == heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		curwidth := 0
		if len(stack) == 0 {
			curwidth = len_height
		} else {
			curwidth = len_height - stack[len(stack)-1] - 1
		}
		max_area = maxvalue(max_area, curwidth*heights[peek_index])
	}
	return max_area
}

func largestRectangleArea1(heights []int) int {
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	var len_hei = len(heights)
	stack := []int{0}
	max_vear := 0
	for i := 0; i < len_hei; i++ {
		for heights[stack[len(stack)-1]] > heights[i] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			max_vear = maxvalue(max_vear, h*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return max_vear
}

func maxvalue(x, y int) int {
	if x < y {
		return y
	}
	return x
}
