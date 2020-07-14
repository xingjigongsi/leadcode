package leadcode

func numDecodings(s string) int {
	if len(s) ==0{
		return 0
	}
	if string(s[0]) == "0"{
		return 0
	}
	dp:=[]int{1,1}
	for i:=1;i<len(s);i++{
		temp:= dp[1]
		if string(s[i]) =="0"{
			if string(s[i-1]) =="1" || string(s[i-1]) =="2"{
				dp[1] = dp[0]
			}else{
				return 0
			}
		}else if string(s[i-1]) =="1" ||
			(string(s[i-1]) =="2" && string(s[i]) >="1" && string(s[i]) <="6"){
			dp[1] = dp[1]+dp[0]
		}
		dp[0] = temp
	}
	return dp[1]
}
