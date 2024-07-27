type RecentCounter struct {
    q []int
}


func Constructor() RecentCounter {
    return RecentCounter{[]int{}}
}


func (c *RecentCounter) Ping(t int) int {
    c.q = append(c.q, t)
    for c.q[0] < t - 3000 {
        c.q = c.q[1:]
    }

    return len(c.q)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */