package tire


type WordDictionary struct {
	isword bool
	next map[string]*WordDictionary
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		isword:false,
		next:make(map[string]*WordDictionary,0),
	}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	for _,v:=range word{
		if _,ok:= this.next[string(v)];!ok{
			this.next[string(v)] = &WordDictionary{
				isword:false,
				next:make(map[string]*WordDictionary,0),
			}
		}
		this = this.next[string(v)]
	}
	if !this.isword{
		this.isword = true
	}
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {

	return this.tree_heper(word,0)
}

func (this *WordDictionary) tree_heper(word string,i int) bool{
	if i == len(word){
		return this.isword
	}
	if word[i]!='.'{
		if _,ok:= this.next[string(word[i])];!ok{
			return false
		}
		return this.next[string(word[i])].tree_heper(word,i+1)
	}else{
		for _,v:=range this.next{
			if v.tree_heper(word,i+1){
				return true
			}
		}
		return false
	}
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */