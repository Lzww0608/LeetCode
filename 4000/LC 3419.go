func minMaxWeight(n int, edges [][]int, threshold int) int {
    g := make([][][2]int, n)
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        g[v] = append(g[v], [2]int{u, w})
    }
    
    h := &hp{{0, 0}}
    heap.Init(h)
    dis := make([]int, n)
    for i := range dis {
        dis[i] = math.MaxInt32
    }
    dis[0] = 0
    for h.Len() > 0 {
        cur := heap.Pop(h).(node)
        u, w := cur.to, cur.weight
        if w > dis[u] {
            continue
        }
        for _, v := range g[u] {
            cur_w := max(w, v[1])
            if dis[v[0]] > cur_w {
                dis[v[0]] = cur_w 
                heap.Push(h, node{v[0], cur_w})
            }
        }
    }
    ans := slices.Max(dis)
    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}

type node struct {
    to, weight int
}
type hp []node
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].weight < h[j].weight}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(node))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}