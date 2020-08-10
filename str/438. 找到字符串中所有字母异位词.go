package str

func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(p)==0{
		return nil
	}
	if len(s)<len(p){
		return nil
	}
	result:=[]int{}
	s_len:= len(s)
	p_len:= len(p)
	i:=0
	for i<s_len && i+p_len<s_len+1{
		s_temp:=s[i:i+p_len]
		nums:=make([]int,26)
		for i:=0;i<len(s_temp);i++ {
			nums[s_temp[i]-'a']++
			nums[p[i]-'a']--
		}
		if checkNum(nums){
			result = append(result,i)
		}
		i++
	}
	return result
}

func checkNum(nums []int) bool{
	for _,v:=range nums {
		if v!=0{
			return false
		}
	}
	return true
}