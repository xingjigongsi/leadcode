package leadcode

/**
题目：
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
**/
var result_s []string

var maps = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	result_s = make([]string, 0)
	letter_helper(digits, 0, "")
	return result_s
}

func letter_helper(digits string, index int, str string) {
	if len(digits) == index {
		result_s = append(result_s, str)
		return
	}
	c := digits[index]
	t := maps[c]
	for _, v := range t {
		letter_helper(digits, index+1, str+string(v))
	}
	return
}
