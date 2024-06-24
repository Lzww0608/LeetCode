func reachableNodes(edges [][]int, maxMoves int, n int) (ans int) {
    g := make([][]neighbor, n)
    for _, e := range edges {
        u, v, cnt := e[0], e[1], e[2]
        g[u] = append(g[u], neighbor{v, cnt + 1})
        g[v] = append(g[v], neighbor{u, cnt + 1})
    }

    dist := dijkstra(g, 0)

    for _, d := range dist {
        if d <= maxMoves {
            ans++
        }
    }

    for _, e := range edges {
        u, v, cnt := e[0], e[1], e[2]
        a := max(maxMoves - dist[u], 0)
        b := max(maxMoves - dist[v], 0)
        ans += min(a + b, cnt)
    }
    return
}

type neighbor struct {
    to, weight int
}

func dijkstra(g [][]neighbor, start int) []int {
    dist := make([]int, len(g))
    for i := range dist {
        dist[i] = math.MaxInt
    }
    dist[start] = 0
    h := hp{{start, 0}}
    for h.Len() > 0 {
        t := heap.Pop(&h).(pair)
        x := t.x
        if dist[x] < t.dist {
            continue
        }
        for _, e := range g[x] {
            y := e.to
            if d := dist[x] + e.weight; d < dist[y] {
                dist[y] = d
                heap.Push(&h, pair{y, d})
            }
        }
    }

    return dist
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