package str

func firstUniqChar(s string) int {
	if s ==""{
		return 0
	}
	s_map:=make(map[int32]int,0)
	for _,v:=range s{
		if _,ok:=s_map[v];!ok{
			s_map[v] = 0
		}else{
			s_map[v] = 1
		}
	}

	for i,v:=range s{
		if s_map[v]==0{
			return i
			break
		}
	}
	return 0
}
