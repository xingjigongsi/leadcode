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


func ladderLength_1(beginWord string, endWord string, wordList []string) int {
	wordset:=make(map[string]string,0)
	for _,v:=range wordList{
		wordset[v] = v
	}
	if _,ok:= wordset[endWord];!ok{
		return 0
	}
	begin_word:= map[string]string{beginWord:beginWord}
	end_word:= map[string]string{endWord:endWord}
	visted:=make(map[string]string,0)
	var c byte
	dep:=1

	for len(begin_word)!=0 && len(end_word)!=0{

		if len(begin_word)>len(end_word){
			begin_word,end_word = end_word,begin_word
		}
		next_wordl:=make(map[string]string,0)
		for _,b_temp:=range begin_word{
			ten:=[]byte(b_temp)
			for j,_:= range ten{
				origan := ten[j]
				for c = 'a';c<='z';c++{
					if c == origan {
						continue
					}
					ten[j] = c
					b_temp_str := string(ten)
					if _,ok:= wordset[b_temp_str];ok{
						if _,ok = end_word[b_temp_str];ok{
							return dep+1
						}
						if _,ok = visted[b_temp_str];!ok{
							next_wordl[b_temp_str] = b_temp_str
							visted[b_temp_str] = b_temp_str
						}
					}
				}
				// 完成后返回原来的值
				ten[j] = origan
			}
		}
		begin_word = next_wordl
		dep++
	}
	return 0
}
