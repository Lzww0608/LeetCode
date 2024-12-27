import (
    "math"
    "container/heap"
)

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
    g := make([][]pair, n)
    rg := make([][]pair, n)
    for _, edge := range edges {
        u, v, d := edge[0], edge[1], edge[2]
        g[u] = append(g[u], pair{v, d})
        rg[v] = append(rg[v], pair{u, d})
    }    

    dis1 := dijkstra(g, src1)
    dis2 := dijkstra(g, src2)
    dis3 := dijkstra(rg, dest)

    var ans int64 = math.MaxInt64
    for i := 0; i < n; i++ {
        if dis1[i] < math.MaxInt64 && dis2[i] < math.MaxInt64 && dis3[i] < math.MaxInt64 {
            ans = min(ans, int64(dis1[i] + dis2[i] + dis3[i]))
        }
    }
    if ans == math.MaxInt64 {
        return -1
    }
    return ans
}

func dijkstra(g [][]pair, start int) []int {
    n := len(g)
    dis := make([]int, n)
    for i := range dis {
        dis[i] = math.MaxInt64
    }
    dis[start] = 0
    h := &hp{{start, 0}}
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        if cur.dis > dis[cur.v] {
            continue
        }

        for _, u := range g[cur.v] {
            if cur.dis + u.dis < dis[u.v] {
                dis[u.v] = cur.dis + u.dis
                heap.Push(h, pair{u.v, dis[u.v]})
            }
        }
    }

    return dis
}

type pair struct {
    v, dis int
}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].dis < h[j].dis}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}