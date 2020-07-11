package leadcode

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1, f2 = f2, f3
	}
	return f3
}

var nums = make(map[int]int)

func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	if re, ok := nums[n]; ok {
		return re
	}
	num := climbStairs1(n-1) + climbStairs1(n-2)
	if _, ok := nums[n]; !ok {
		nums[n] = num
	}
	return num
}

func climbStairs3(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs3(n-1) + climbStairs3(n-2)
}

func climbStairs4(n int) int {
	nums := make([]int, n)
	nums[0] = 1
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums[n]
}

//func main(){
//	result := climbStairs1(9)
//	fmt.Println(result)
//
//}
