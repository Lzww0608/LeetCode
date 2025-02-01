
import (
	"container/heap"
	"math"
)
func minimumCost(n int, highways [][]int, discounts int) int {
    g := make([][][2]int, n)
    for _, v := range highways {
        a, b, w := v[0], v[1], v[2]
        g[a] = append(g[a], [2]int{b, w})
        g[b] = append(g[b], [2]int{a, w})
    }

    dis := make([][]int, n)
    for i := range dis {
        dis[i] = make([]int, discounts + 1)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt32
        }
    }
    for j := range dis[0] {
        dis[0][j] = 0
    }

    h := &hp{{0, 0, discounts}}
    ans := math.MaxInt32
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        u, w, d := cur.to, cur.weight, cur.dsc
        if u == n - 1 {
            ans = min(ans, w)
        }
        for _, v := range g[u] {
            x := v[0]
            cur_w := v[1] + w 
            if dis[x][d] > cur_w {
                dis[x][d] = cur_w
                heap.Push(h, pair{x, cur_w, d})
            }

            if d > 0 {
                cur_w = v[1] / 2 + w
                if dis[x][d-1] > cur_w {
                    dis[x][d-1] = cur_w
                    heap.Push(h, pair{x, cur_w, d - 1})
                }
            }
        }
    }

    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}

type pair struct {
    to, weight, dsc int
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

