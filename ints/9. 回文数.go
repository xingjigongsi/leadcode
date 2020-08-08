package ints


func isPalindrome(x int) bool {
	if x<0 || (x%10==0 && x!=0){
		return false
	}
	new_x:=0
	if x>new_x{
		new_x = new_x*10+x%10
		x = x/10
	}
	return x == new_x || x == new_x/10
}


func isPalindrome_1(x int) bool {
	if x<0{
		return false
	}
	new_x:=0
	item:=x
	for x>0{
		new_x = new_x*10+x%10
		x = x/10
	}
	return new_x==item
}
