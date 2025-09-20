type pair struct {
    s, d, t int
}

type Router struct {
    p map[pair]bool
    q []pair
    m map[int][]int
    n int
}


func Constructor(memoryLimit int) Router {
    p := make(map[pair]bool)
    q := []pair{}
    n := memoryLimit
    m := make(map[int][]int)
    return Router{p, q, m, n}
}


func (c *Router) AddPacket(source int, destination int, timestamp int) bool {
    tmp := pair{source, destination, timestamp}
    if c.p[tmp] {
        return false
    } 
    c.p[tmp] = true 
    c.q = append(c.q, tmp)
    c.m[destination] = append(c.m[destination], timestamp)
    if len(c.q) > c.n {
        tmp = c.q[0]
        delete(c.p, tmp)
        c.q = c.q[1:]
        c.m[tmp.d] = c.m[tmp.d][1:]
    }
    return true
}


func (c *Router) ForwardPacket() []int {
    if len(c.q) == 0 {
        return nil
    }

    ans := c.q[0]
    c.q = c.q[1:]
    delete(c.p, ans)
    c.m[ans.d] = c.m[ans.d][1:]
    return []int{ans.s, ans.d, ans.t}
}


func (c *Router) GetCount(destination int, startTime int, endTime int) int {
    l := sort.Search(len(c.m[destination]), func(i int) bool {
        return c.m[destination][i] >= startTime
    })
    r := sort.Search(len(c.m[destination]), func(i int) bool {
        return c.m[destination][i] > endTime
    })

    return r - l
}


/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */