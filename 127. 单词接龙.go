package leadcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordset:=make(map[string]string,0)
	for _,v:=range wordList{
		wordset[v] = v
	}
	if _,ok:= wordset[endWord];!ok{
		return 0
	}
	queue:= []string{beginWord}
	// 访问过的单词
	vistedSet:=make(map[string]string,0)
	var c byte
	dep:= 1
	for len(queue)>0{
		len_q:= len(queue)
		for i:=0;i<len_q;i++{
			q:= []byte(queue[0])
			queue = queue[1:]
			for j,str:=range q{
				origin_c:= q[j]
				for c='a';c<='z';c++{
					if str == c {
						continue
					}
					q[j] = c
					q_str:=string(q)
					if _,ok:=wordset[q_str];ok{
						if q_str == endWord{
							return dep+1
						}
						if _,ok= vistedSet[q_str];!ok{
							vistedSet[q_str] = q_str
							queue = append(queue,q_str)
						}
					}
				}
				q[j] = origin_c
			}

		}
		dep++
	}
	return 0
}




