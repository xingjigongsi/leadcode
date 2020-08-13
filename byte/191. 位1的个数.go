package byte

func hammingWeight(num uint32) int {
	result,power:=0,uint32(31)
	for num>0{
		temp:= num&1
		if temp==1{
			result++
		}
		num = num>>1
		power--
	}
	return result
}
