func shortestPath(n int, edges [][]int, labels string, k int) int {
    g := make([][]pair, n)
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        g[u] = append(g[u], pair{v, w})
    }

    dis := make([]int, n * k)
    for i := range dis {
        dis[i] = math.MaxInt32
    }

    idx := func(u, cnt int) int {
        return u * k + cnt - 1
    }

    h := &hp{}
    heap.Push(h, info{0, 0, 1})
    for h.Len() > 0 {
        cur := heap.Pop(h).(info)
        if cur.w > dis[idx(cur.u, cur.cnt)] {
            continue
        } 
        
        if cur.u == n - 1 {
            return cur.w
        }

        for _, e := range g[cur.u] {
            next_cnt := 1
            if labels[e.u] == labels[cur.u] {
                if next_cnt = cur.cnt + 1; next_cnt > k {
                    continue
                }
            } 

            next := idx(e.u, next_cnt)
            if dis[next] > cur.w + e.w {
                dis[next] = cur.w + e.w 
                heap.Push(h, info{e.u, cur.w + e.w, next_cnt})
            } 
        } 
    }

    return -1
}

type pair struct {
    u, w int 
}

type info struct {
    u, w, cnt int
}

type hp []info 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].w < h[j].w}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(info))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return
}