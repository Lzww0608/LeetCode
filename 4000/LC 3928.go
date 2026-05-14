func minCost(n int, prices []int, roads [][]int) []int {
    g := make([][]pair, n)
    for _, road := range roads {
        u, v, c, t := road[0], road[1], road[2], road[3]
        g[u] = append(g[u], pair{v, c, t})
        g[v] = append(g[v], pair{u, c, t})
    }
    
    dij := func(src, cost int) []int {
        dis := make([]int, n)
        for i := range dis {
            dis[i] = math.MaxInt / 3 
        }
        dis[src] = 0 

        h := &hp{}
        heap.Push(h, pair{src, 0, 0})
        for h.Len() > 0 {
            cur := heap.Pop(h).(pair)
            if dis[cur.v] < cur.w {
                continue
            }
            for _, v := range g[cur.v] {
                w := cur.w + v.w
                if cost == 2 {
                    w = cur.w + v.w * v.t
                } else if w >= prices[src] {
                    continue
                }
                if dis[v.v] > w {
                    dis[v.v] = w 
                    heap.Push(h, pair{v.v, w, 0})
                }
            }
        }

        return dis
    }

    dis := make([][]int, n)
    re_dis := make([][]int, n)
    for i := range n {
        dis[i] = dij(i, 1)
        re_dis[i] = dij(i, 2)
    }

    ans := make([]int, n)
    for i := range n {
        ans[i] = prices[i]
        for j := range n {
            if j == i {
                continue
            }

            ans[i] = min(ans[i], dis[i][j] + re_dis[j][i] + prices[j])
        }
    }

    return ans
}

type pair struct {
    v, w, t int 
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