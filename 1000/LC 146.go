type node struct {
    k, v int
}

type LRUCache struct {
    cap int
    list *list.List
    keyToNode map[int]*list.Element
}


func Constructor(capacity int) LRUCache {
    return LRUCache{capacity, list.New(), map[int]*list.Element{}}
}


func (this *LRUCache) Get(key int) int {
    res := this.keyToNode[key]
    if res == nil {
        return -1
    }
    this.list.MoveToFront(res)
    return res.Value.(node).v
}


func (this *LRUCache) Put(key int, value int)  {
    if res := this.keyToNode[key]; res != nil {
        res.Value = node{key, value}
        this.list.MoveToFront(res)
        return
    }
    newNode := node{key, value}
    this.keyToNode[key] = this.list.PushFront(newNode)
    if len(this.keyToNode) > this.cap {
        delete(this.keyToNode, this.list.Remove(this.list.Back()).(node).k)
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
