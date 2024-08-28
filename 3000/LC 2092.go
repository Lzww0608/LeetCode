type Edge struct {
    edge int
    weight int
}

type hp []Edge
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].weight < h[j].weight}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(Edge))
} 
func (h *hp) Pop() (x any) {
    old := *h
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}

func findAllPeople(n int, meetings [][]int, firstPerson int) (ans []int) {
    g := make([][]Edge, n)
    for _, e := range meetings {
        u, v, w := e[0], e[1], e[2]
        g[u] = append(g[u], Edge{v, w})
        g[v] = append(g[v], Edge{u, w})
    } 

    dis := make([]int, n)
    for i := range dis {
        dis[i] = 1 << 31 - 1
    }
    dis[0], dis[firstPerson] = 0, 0
    h := &hp{Edge{0, 0}, Edge{firstPerson, 0}}
    heap.Init(h)
    vis := make([]bool, n)

    for h.Len() > 0 {
        cur := heap.Pop(h).(Edge)
        u := cur.edge
        if vis[u] {
            continue
        }
        vis[u] = true
        ans = append(ans, u)
        for _, e := range g[u] {
            v, w := e.edge, e.weight
            if vis[v] || w < cur.weight {
                continue
            }
            heap.Push(h, e)
        }
    }

    return 
}