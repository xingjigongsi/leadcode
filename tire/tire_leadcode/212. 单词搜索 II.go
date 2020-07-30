package tire_leadcode

var w,h int
var result []string
var targets = [][]int{{-1,0},{0,-1},{1,0},{0,1}}  // 方向
var vistied [][]bool
var set_result map[string]bool

func findWords(board [][]byte, words []string) []string {
	if len(board) ==0 || len(words)==0{
		return []string{}
	}
	result = make([]string,0)
	tire:= Constructor()
	for _,v:=range words{
		tire.Insert(v)
	}
	w = len(board)
	h = len(board[0])
	vistied = make([][]bool,w)
	set_result = make(map[string]bool,0)
	for i,_:=range vistied{
		vistied[i] = make([]bool,h)
	}
	for i:=0;i<w;i++{
		for j:=0;j<h;j++{
			findWords_helper(board,tire,i,j,"")
		}
	}
	return result
}

func checkwidth(x,y int) bool{
	return x>=0 && x<w && y>=0 && y<h
}

func findWords_helper(board [][]byte,words_tree *Trie,i,j int,cur_word string){
	cur_word += string(board[i][j])
	if v,ok:=words_tree.next[string(board[i][j])];!ok{
		return
	}else{
		if  _,ok:=set_result[cur_word];!ok && v.isword {
			result = append(result,cur_word)
			set_result[cur_word] = true
		}
		words_tree = words_tree.next[string(board[i][j])]
	}
	vistied[i][j] =true
	for _,m:=range targets{
		new_i:=m[0]+i
		new_j:=m[1]+j
		if checkwidth(new_i,new_j) && !vistied[new_i][new_j]{
			findWords_helper(board,words_tree,new_i,new_j,cur_word)
		}
	}
	vistied[i][j] = false
}




// 前缀树
type Trie struct {
	isword bool
	next map[string]*Trie
}


/** Initialize your data structure here. */
func Constructor() *Trie {
	return &Trie{
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
