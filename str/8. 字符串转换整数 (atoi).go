package str

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(str string) int {
	sign,res,i:=1,0,0
	if len(str)==0{
		return 0
	}
	str = strings.TrimSpace(str)
	for i<len(str){
		if i>=1 && (str[i]=='-' || str[i]=='+'){
			break
		}
		if str[i]=='-'{
			sign = -1
			i++
			continue
		}
		if str[i]=='+'{
			sign = 1
			i++
			continue
		}
		if !unicode.IsDigit(rune(str[i])){
			break
		}
		r:=str[i]-'0'
		res = res*10+int(r)
		if res>=math.MaxInt32 && sign==1{
			return  math.MaxInt32
		}
		if res>math.MaxInt32 && sign==-1{
			return math.MinInt32
		}
		i++
	}

	if sign==1{
		return res
	}
	return -res
}
