type MapSum struct {
    cnt int
    children [26]*MapSum
}

var root *MapSum
var m map[string]int

/** Initialize your data structure here. */
func Constructor() MapSum {
    var cnt int
    var children [26]*MapSum
    root = &MapSum{cnt, children}
    m = map[string]int{}
    return *root
}


func (this *MapSum) Insert(key string, val int) {
    t := m[key]
    m[key] = val
    cur := root
    for i := range key {
        idx := int(key[i] - 'a')
        if cur.children[idx] == nil {
            cur.children[idx] = &MapSum{}
        }
        cur = cur.children[idx]
        cur.cnt += val - t
    }
}


func (this *MapSum) Sum(prefix string) int {
    cur := root
    for i := range prefix {
        idx := int(prefix[i] - 'a')
        if cur.children[idx] == nil {
            return 0
        }
        cur = cur.children[idx]
    }
    return cur.cnt
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
