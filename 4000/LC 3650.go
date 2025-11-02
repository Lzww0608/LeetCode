func minCost(n int, edges [][]int) int {
    g := make([][]pair, n)
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        g[u] = append(g[u], pair{v, w})
        g[v] = append(g[v], pair{u, w * 2})
    }

    dis := make([]int, n)
    for i := range n {
        dis[i] = math.MaxInt32
    }
    dis[0] = 0

    h := &hp{}
    heap.Push(h, pair{0, 0})
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        if cur.w > dis[cur.u] {
            continue
        }

        for _, v := range g[cur.u] {
            if cur.w + v.w < dis[v.u] {
                dis[v.u] = cur.w + v.w
                heap.Push(h, pair{v.u, dis[v.u]})
            }
        }
    }

    if dis[n - 1] >= math.MaxInt32 {
        return -1
    }
    return dis[n - 1]
}

type pair struct {
    u, w int 
}

type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].w < h[j].w}
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