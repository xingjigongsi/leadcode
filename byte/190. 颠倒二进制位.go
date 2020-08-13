package byte


func reverseBits(num uint32) uint32 {
	ret,n:= uint32(0),uint32(31)
	for num>0{
		ret+=(num & 1)<<n  // 获取最后一位
		num = num>>1
		n--
	}
	return ret;
}
