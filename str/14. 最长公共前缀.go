package str

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	for i:=0;i<len(strs[0]);i++{
		c:=strs[0][i]
		for j:=1;j<len(strs);j++{
			if len(strs[j])==i || strs[j][i]!=c{
				return  strs[0][0:i]
			}
		}
	}
	return strs[0]
}
