
import "math"
func minimumDistance(n int, edges [][]int, s int, marked []int) int {
    g := make([][][2]int, n)
    for _, v := range edges {
        a, b, w := v[0], v[1], v[2]
        g[a] = append(g[a], [2]int{b, w})
    } 
    dis := make([]int, n)
    for i := range dis {
        dis[i] = math.MaxInt32
    }
    dis[s] = 0 

    h := &hp{{s, 0}}
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        u, w := cur.to, cur.weight
        if dis[u] < w {
            continue
        }
        for _, v := range g[u] {
            cur_w := v[1] + w 
            if cur_w < dis[v[0]] {
                dis[v[0]] = cur_w
                heap.Push(h, pair{v[0], cur_w})
            }
        }
    }

    ans := math.MaxInt32
    for _, i := range marked {
        ans = min(ans, dis[i])
    }
    if ans == math.MaxInt32 {
        return -1
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