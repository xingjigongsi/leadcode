package leadcode


func mySqrt(x int) int {
	if x==0 || x==1{
		return x
	}
	left:=1
	right:=x
	for left<=right{
		midle:= (left+right)/2
		if midle*midle>x{
			right = midle-1
		}else{
			left = midle+1
		}
	}
	return right
}
