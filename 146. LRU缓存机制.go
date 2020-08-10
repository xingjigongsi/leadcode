package leadcode

import (
	"container/list"
	"fmt"
)


type LRUCache struct {
	list  *list.List
	value int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{list:list.New()}
}


func (this *LRUCache) Get(key int) int {
	for i := this.list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}


func (this *LRUCache) Put(key int, value int)  {

}