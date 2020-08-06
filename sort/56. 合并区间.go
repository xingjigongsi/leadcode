package sort

import "sort"

/*
示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0{
		return [][]int{}
	}
	//排序
	sort.Slice(intervals, func(i,j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	result:=[][]int{}
	left:= intervals[0][0]
	right:= intervals[0][1]
	for i:=1;i<len(intervals);i++{
		if intervals[i][0]<=right{
			if intervals[i][1]>right{
				right = intervals[i][1]
			}

		}else{
			result = append(result,[]int{left,right})
			left = intervals[i][0]
			right =intervals[i][1]
		}

	}
	result = append(result,[]int{left,right})
	return result
}
