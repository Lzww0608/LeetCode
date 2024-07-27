func networkBecomesIdle(edges [][]int, patience []int) (ans int) {
    n := len(patience)
    g := make([][]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    dist := make([]int, n)
    for i := range dist {
        dist[i] = math.MaxInt32
    }
    dist[0] = 0
    h := hp{pair{0, 0}}
    heap.Init(&h)

    for h.Len() > 0 {
        tmp := heap.Pop(&h).(pair)
        d, u := tmp.dis, tmp.x

        if dist[u] < d {
            continue
        }
        for _, v := range g[u] {
            if dist[u]+1 < dist[v] {
                dist[v] = dist[u] + 1
                heap.Push(&h, pair{v, dist[v]})
            }
        }
    }

    for i, x := range patience[1:] {
        roundTripTime := dist[i+1] * 2
        lastResendTime := (roundTripTime - 1) / x * x
        totalTime := roundTripTime + lastResendTime
        ans = max(ans, totalTime + 1)
    }

    return
}

type pair struct {
    x, dis int
}
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
