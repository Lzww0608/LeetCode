type entry struct {
    key, value, freq int
}

type LFUCache struct {
    capacity int
    minFreq int
    keyToNode map[int]*list.Element
    freqToList map[int]*list.List
}




func Constructor(capacity int) LFUCache {
    return LFUCache {
        capacity: capacity,
        keyToNode: map[int]*list.Element{},
        freqToList: map[int]*list.List{},
    }
}

func (this *LFUCache) pushFront(e *entry) {
    if _, ok := this.freqToList[e.freq]; !ok {
        this.freqToList[e.freq] = list.New()
    }
    this.keyToNode[e.key] = this.freqToList[e.freq].PushFront(e);
}

func (this *LFUCache) getEntry(key int) *entry {
    node := this.keyToNode[key]
    if node == nil {
        return nil
    }
    e := node.Value.(*entry)
    l := this.freqToList[e.freq]
    l.Remove(node)
    if l.Len() == 0 {
        delete(this.freqToList, e.freq)
        if this.minFreq == e.freq {
            this.minFreq++
        }
    }
    e.freq++
    this.pushFront(e)
    return e
}


func (this *LFUCache) Get(key int) int {
    if e := this.getEntry(key); e != nil {
        return e.value
    }
    return -1
}


func (this *LFUCache) Put(key int, value int)  {
    if e := this.getEntry(key); e != nil {
        e.value = value
        return
    }
    if len(this.keyToNode) == this.capacity {
        l := this.freqToList[this.minFreq]
        delete(this.keyToNode, l.Remove(l.Back()).(*entry).key)
        if l.Len() == 0 {
            delete(this.freqToList, this.minFreq)
        }
        
    }
    this.pushFront(&entry{key, value, 1})
        this.minFreq = 1
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
