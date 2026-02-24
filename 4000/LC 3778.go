func minCostExcludingMax(n int, edges [][]int) int64 {
    g := make([][]pair, n)
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        g[u] = append(g[u], pair{v, w})
        g[v] = append(g[v], pair{u, w})
    }

    dis := make([][2]int, n)
    for i := range dis {
        dis[i] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
    }
    dis[0] = [2]int{0, 0}
    h := &hp{{0, 0, 0}}
    for h.Len() > 0 {
        cur := heap.Pop(h).(tuple)
        if cur.mx == -1 {
            if cur.d > dis[cur.u][1] {
                continue
            }
            for _, v := range g[cur.u] {
                d := cur.d + v.w 
                if d < dis[v.v][1] {
                    dis[v.v][1] = d 
                    heap.Push(h, tuple{v.v, d, -1})
                }
            }
        } else {
            if cur.d <= dis[cur.u][0] {
                for _, v := range g[cur.u] {
                    d := cur.d + v.w 
                    if d < dis[v.v][0] {
                        dis[v.v][0] = d 
                        heap.Push(h, tuple{v.v, d, max(cur.mx, v.w)})
                    }
                }
            } 

            for _, v := range g[cur.u] {
                if cur.d < dis[v.v][1] {
                    dis[v.v][1] = cur.d
                    heap.Push(h, tuple{v.v, cur.d, -1})
                }
            }
        }
    }

    return int64(min(dis[n - 1][0], dis[n - 1][1]))
}

type pair struct {
    v, w int
}
type tuple struct {
    u, d, mx int
}
type hp []tuple 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].d < h[j].d}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(tuple))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return
}
