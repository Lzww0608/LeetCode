
import (
	"container/heap"
	"math"
)
func findAnswer(n int, edges [][]int) []bool {
    ans := make([]bool, len(edges))
    g := make([][][]int, n)
    dis := make([]int, n)
    for i := range dis {
        dis[i] = math.MaxInt32
    }   
    dis[0] = 0
    id := make(map[[2]int]int)
    for i, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        g[u] = append(g[u], []int{v, w})
        g[v] = append(g[v], []int{u, w})
        id[[2]int{u, v}] = i 
        id[[2]int{v, u}] = i
    }
    h := &hp{pair{0, 0}}
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        for _, v := range g[cur.to] {
            w := cur.weight + v[1]
            if w < dis[v[0]] {
                dis[v[0]] = w
                heap.Push(h, pair{v[0], dis[v[0]]})
            }
        }
    }
    if dis[n-1] == math.MaxInt32 {
        return ans
    }

    q := [][]int{{n - 1, dis[n-1]}}
    vis := make([]bool, n)
    vis[n-1] = true
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        for _, v := range g[cur[0]] {
            if dis[v[0]] == cur[1] - v[1] {
                ans[id[[2]int{cur[0], v[0]}]] = true
                if !vis[v[0]] {
                    vis[v[0]] = true
                    q = append(q, []int{v[0], dis[v[0]]})
                }
                
            }
        }
    }

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