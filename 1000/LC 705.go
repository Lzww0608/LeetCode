type MyHashSet struct {
    bucket [40000]int
}


func Constructor() MyHashSet {
    m := [40000]int{}
    return MyHashSet{m}
}

func (this *MyHashSet) Get(idx, offset int) int {
    return (this.bucket[idx] >> offset) & 1 
}

func (this *MyHashSet) Set(idx, offset int, f bool) {
    if f {
        this.bucket[idx] |= 1 << offset
    } else {
        this.bucket[idx] &= ^(1 << offset)
    }
}

func (this *MyHashSet) Add(key int)  {
    idx, offset := key / 32, key % 32
    this.Set(idx, offset, true)
}


func (this *MyHashSet) Remove(key int)  {
    idx, offset := key / 32, key % 32
    this.Set(idx, offset, false)
}


func (this *MyHashSet) Contains(key int) bool {
    return this.Get(key / 32, key % 32) == 1
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
