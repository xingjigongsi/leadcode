package leadcode

// 时间复杂度 O(mk)
// 空间复杂度 O(k)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	cont:=make([]int,0)
	for i:=0;i<len(nums);i++{
		for _,v:=range cont{
			if checknums(nums[i],v,t){
				return true
			}
		}
		cont = append(cont,nums[i])
		if len(cont) == k+1{
			cont = cont[1:]
		}
	}
	return false
}

func checknums(x,y,t int) bool{
	res:=x-y
	if res<0{
		res = -res
	}
	return res<=t

}
