func minTime(n int, edges [][]int) int {
    g := make([][][]int, n)
    for _, edge := range edges {
        g[edge[0]] = append(g[edge[0]], edge)
    }

    dis := make([]int, n)
    for i := range dis {
        dis[i] = math.MaxInt32
    }
    dis[0] = 0

    h := &hp{}
    heap.Push(h, pair{0, 0})
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        if cur.t > dis[cur.u] {
            continue
        }
        for _, v := range g[cur.u] {
            if cur.t > v[3] {
                continue
            }

            t := max(cur.t, v[2]) + 1
            if t < dis[v[1]] {
                dis[v[1]] = t 
                heap.Push(h, pair{v[1], t})
            }
        }
    }

    if dis[n - 1] >= math.MaxInt32 {
        return -1
    }
    return dis[n - 1]
}

type pair struct {
    u, t int 
}

type hp []pair 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].t < h[j].t}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return
}