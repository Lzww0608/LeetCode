
import (
	"container/heap"
	"math"
)
func shortestPathWithHops(n int, edges [][]int, s int, d int, k int) int {
    g := make([][][2]int, n)
    for _, v := range edges {
        a, b, w := v[0], v[1], v[2]
        g[a] = append(g[a], [2]int{b, w})
        g[b] = append(g[b], [2]int{a, w})
    }
    dis := make([][]int, n)
    for i := range dis {
        dis[i] = make([]int, k + 1)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt32
        }
    }
    
    dis[s][k] = 0
    h := &hp{{s, 0, k}}
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        u, w, t := cur.to, cur.weight, cur.k
        if dis[u][t] < w {
            continue
        }
        for _, v := range g[u] {
            cur_w := w + v[1]
            if dis[v[0]][t] > cur_w {
                dis[v[0]][t] = cur_w
                heap.Push(h, pair{v[0], cur_w, t})
            }

            if t > 0 && dis[v[0]][t-1] > w {
                dis[v[0]][t-1] = w
                heap.Push(h, pair{v[0], w, t - 1})
            }
        }
    }

    return slices.Min(dis[d])
}

type pair struct {
    to, weight, k int
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
