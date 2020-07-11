package leadcode

/**
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/
func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}
	m := len(text1)
	n := len(text2)
	resu := make([][]int, m+1)
	for i, _ := range resu {
		resu[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				resu[i][j] = resu[i-1][j-1] + 1
			} else {
				resu[i][j] = max(resu[i-1][j], resu[i][j-1])
			}
		}
	}
	return resu[m][n]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

var longmaps map[int]map[int]int

func longestCommonSubsequence_1(text1 string, text2 string) int {
	longmaps = make(map[int]map[int]int, 0)
	m := len(text1) - 1
	n := len(text2) - 1
	return longest_helper(text1, text2, m, n)

}

func longest_helper(text1 string, text2 string, i, j int) int {
	if i == -1 || j == -1 {
		return 0
	}
	if v, ok := longmaps[i][j]; ok {
		return v
	}
	var t int
	if text1[i] == text2[j] {
		t = longest_helper(text1, text2, i-1, j-1) + 1
	} else {
		t = max(longest_helper(text1, text2, i-1, j), longest_helper(text1, text2, i, j-1))
	}
	if _, ok := longmaps[i][j]; !ok {
		longmaps[i][j] = t
	}
	return t

}
