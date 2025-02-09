
import (
	"container/heap"
	"math"
)
func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
    n := len(charge)
    g := make([][][2]int, n)
    for _, v := range paths {
        a, b, w := v[0], v[1], v[2]
        g[a] = append(g[a], [2]int{b, w})
        g[b] = append(g[b], [2]int{a, w})
    }

    dis := make([][]int, n)
    for i := range dis {
        dis[i] = make([]int, cnt + 1)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt32
        }
    }

    h := &hp{{start, 0, 0}}
    heap.Init(h)
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        u, w, t := cur.to, cur.weight, cur.cnt
        if w > dis[u][t] {
            continue
        } else if u == end {
            return w
        }
        if t < cnt {
            if w + charge[u] < dis[u][t+1] {
                dis[u][t+1] = w + charge[u]
                heap.Push(h, pair{u, dis[u][t+1], t + 1})
            }
        }

        for _, v := range g[u] {
            if t < v[1] {
                continue
            }
            cur_w := w + v[1]
            if cur_w < dis[v[0]][t-v[1]] {
                dis[v[0]][t-v[1]] = cur_w
                heap.Push(h, pair{v[0], cur_w, t - v[1]})
            }
        }
    }

    return -1
}

type pair struct {
    to, weight, cnt int
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