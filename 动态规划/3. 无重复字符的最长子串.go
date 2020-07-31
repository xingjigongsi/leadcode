package leadcode

/**
"abcabcbb"
 */
func lengthOfLongestSubstring(s string) int {
	longer_map:=make(map[byte]int,0)
	max:=0
	start:=0
	for i:=0;i<len(s);i++{
		c:=s[i]
		if v,ok:=longer_map[c];ok{
			start = max_string(start,v+1)  //多种情况选取最大的
		}
		longer_map[c] = i
		max = max_string(max,i-start+1)
	}
	return max
}

func max_string(x,y int) int{
	if x>y{
		return x
	}
	return y
}