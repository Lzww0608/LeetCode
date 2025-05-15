type MRUQueue struct {
    a []int
    v []int
}

func (c *MRUQueue) update(i int) {
    for ; i < len(c.a); i += i & (-i) {
        c.a[i]++
    }
}

func (c *MRUQueue) query(i int) (res int) {
    for ; i > 0; i -= i & (-i) {
        res += c.a[i]
    }
    return
}


func Constructor(n int) MRUQueue {
    a := make([]int, n + 2000)
    v := make([]int, n)
    for i := range v {
        v[i] = i + 1
    }

    return MRUQueue{a, v}
}


func (c *MRUQueue) Fetch(k int) int {
    l, r := 1, len(c.v)
    for l < r {
        mid := l + ((r - l) >> 1)
        if mid - c.query(mid) >= k {
            r = mid
        } else {
            l = mid + 1
        }
    }

    c.v = append(c.v, c.v[l-1])
    c.update(l)
    return c.v[l-1]
}


/**
 * Your MRUQueue object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Fetch(k);
 */