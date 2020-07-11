package leadcode

import "sort"

/**
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。对每个孩子 i ，都有一个胃口值 gi ，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j ，都有一个尺寸 sj 。如果 sj >= gi ，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/assign-cookies
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/
func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	result := 0
	sort.Sort(sort.Reverse(sort.IntSlice(g)))
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	g_i, s_i := 0, 0
	for g_i < len(g) && s_i < len(s) {
		if g[g_i] <= s[s_i] {
			result++
			g_i++
			s_i++
		} else {
			g_i++
		}
	}
	return result
}
