package str

import (
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	if len(s) == 0{
		return ""
	}
	spaceRe, _ := regexp.Compile(`\s+`)
	f:= spaceRe.Split(strings.Trim(s," "),-1)
	i,j:=0,len(f)-1
	for i<j{
		f[i],f[j] = f[j],f[i]
		i++
		j--
	}
	result:= strings.Join(f," ")
	return result
}
