package tire


type Trie struct {
	isword bool
	next map[string]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		isword:false,
		next:make(map[string]*Trie,0),
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {

	for _,v:=range word{
		if _,ok:=this.next[string(v)];!ok{
			this.next[string(v)] = &Trie{
				isword:false,next:make(map[string]*Trie,0),
			}
		}
		this,_ = this.next[string(v)]
	}
	if this.isword==false{
		this.isword = true
	}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _,v:=range word{
		if _,ok:=this.next[string(v)];!ok{
			return false
		}
		this,_ = this.next[string(v)]
	}
	return this.isword
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _,v:=range prefix{
		if _,ok:=this.next[string(v)];!ok{
			return false
		}
		this,_ = this.next[string(v)]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
