
import "math"
func minCost(n int, roads [][]int, appleCost []int, k int) []int64 {
    g := make([][][2]int, n + 1)
    for _, v := range roads {
        a, b, c := v[0] - 1, v[1] - 1, v[2]
        g[a] = append(g[a], [2]int{b, c * (k + 1)})
        g[b] = append(g[b], [2]int{a, c * (k + 1)})
    }
    ans := make([]int64, n)
    for i := 0; i < n; i++ {
        ans[i] = int64(appleCost[i])
        g[n] = append(g[n], [2]int{i, appleCost[i]})  
    }

    dijkstra := func(start int) {
        h := &hp{{start, 0}}
        heap.Init(h)
        dis := make([]int, n + 1)
        for i := 0; i < n; i++ {
            dis[i] = math.MaxInt
        }
        dis[start] = 0
        for h.Len() > 0 {
            cur := heap.Pop(h).(pair)
            u, w := cur.to, cur.weight
            if dis[u] < w {
                continue
            }
            for _, v := range g[u] {
                if v[1] + w < dis[v[0]] {
                    dis[v[0]] = v[1] + w
                    ans[v[0]] = min(ans[v[0]], int64(v[1] + w))
                    heap.Push(h, pair{v[0], v[1] + w}) 
                }
            }
        }
    }
    dijkstra(n)

    return ans
}

type pair struct {
    to, weight int
}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].weight < h[j].weight}
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