func minTimeMaxPower(n int, edges [][]int, power int, cost []int, source int, target int) []int64 {
    g := make([][]pair, n)
    for _, v := range edges {
        u, v, t := v[0], v[1], v[2]
        g[u] = append(g[u], pair{v, t, 0})
    }    

    h := &hp{}
    heap.Push(h, pair{source, 0, power})
    dis := make([][]int, n)
    for i := range dis {
        dis[i] = make([]int, power + 1)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt / 2
        }
    }
    dis[source][power] = 0

    ans := []int64{math.MaxInt64 / 2, math.MaxInt64 / 2}
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        if cur.t > dis[cur.u][cur.p] {
            continue
        } else if cur.u == target {
            if int64(cur.t) < ans[0] {
                ans[0], ans[1] = int64(cur.t), int64(cur.p) 
            } else if int64(cur.t) == ans[0] {
                ans[1] = max(ans[1], int64(cur.p))
            }
            continue
        }

        for _, v := range g[cur.u] {
            nxt, t, p := v.u, cur.t + v.t, cur.p - cost[cur.u]
            if p >= 0 && dis[nxt][p] > t {
                dis[nxt][p] = t 
                heap.Push(h, pair{nxt, t, p})
            }
        }
    }

    if ans[0] == math.MaxInt / 2 {
        return []int64{-1, -1}
    }
    return ans
}


type pair struct {
    u, t, p int 
}

type hp []pair 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {
    if h[i].t == h[j].t {
        return h[i].p > h[j].p
    }

    return h[i].t < h[j].t
}
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