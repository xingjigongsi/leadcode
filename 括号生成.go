package leadcode

var results_s []string

func generateParenthesis(n int) []string {
	results_s = make([]string, 0)
	getdata(0, 0, n, "")
	return results_s
}

func getdata(left, right, n int, str string) {
	if left == n && right == n {
		results_s = append(results_s, str)
		return
	}
	if left <= 3 {
		getdata(left+1, 0, n, str+"(")
	}
	if left >= right {
		getdata(left+1, right+1, n, str+")")
	}
}
