const N int = 26

type MapSum struct {
    cnt int
    children [N]*MapSum
}

var root *MapSum
//var m map[string]int

/** Initialize your data structure here. */
func Constructor() MapSum {
    var cnt int
    var children [N]*MapSum
    root = &MapSum{cnt, children}
    //m = map[string]int{}
    return *root
}


func (this *MapSum) Insert(key string, val int) {
    cur := root
    for i := range key {
        idx := int(key[i] - 'a')
        if cur.children[idx] == nil {
            cur.children[idx] = &MapSum{}
        }
        cur = cur.children[idx]
    }
    cur.cnt = val
}


func (this *MapSum) Sum(prefix string) (ans int) {
    cur := root
    for i := range prefix {
        idx := int(prefix[i] - 'a')
        if cur.children[idx] == nil {
            return 0
        }
        cur = cur.children[idx]
    }

    q := []*MapSum{cur}
    for len(q) > 0 {
        cur = q[0]
        ans += cur.cnt
        q = q[1:]
        for i := 0; i < N; i++ {
            if cur.children[i] != nil {
                q = append(q, cur.children[i])
            }
        }
    }

    return
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
