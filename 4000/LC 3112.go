func minimumTime(n int, edges [][]int, disappear []int) []int {
    g := make([][]neighbor, n)
    for _, e := range edges {
        a, b, c := e[0], e[1], e[2]
        g[a] = append(g[a], neighbor{b, c})
        g[b] = append(g[b], neighbor{a, c})
    }

    dijkstra := func() []int {
        dist := make([]int, len(g))
        for i := range dist {
            dist[i] = -1
        }
        dist[0] = 0
        h := hp{{0, 0}}
        for h.Len() > 0 {
            t := heap.Pop(&h).(pair)
            x := t.x
            if dist[x] < t.dist {
                continue
            }
            for _, e := range g[x] {
                y := e.to
                if d := dist[x] + e.weight; d < disappear[y] && (dist[y] < 0 || d < dist[y])  {
                    dist[y] = d
                    heap.Push(&h, pair{y, d})
                }
            }
        }

        return dist
    }

    return dijkstra()
}



type neighbor struct {
    to, weight int
}


type pair struct {
    x, dist int
}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].dist < h[j].dist}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() (x any) {
    old := *h
    n := len(old)
    x, *h = old[n-1], old[:n-1]
    return
}