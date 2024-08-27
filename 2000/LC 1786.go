const MOD int = 1_000_000_007
type Edge struct {
    node int
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

func countRestrictedPaths(n int, edges [][]int) int {
    g := make([][]Edge, n)
    for _, e := range edges {
        u, v, w := e[0] - 1, e[1] - 1, e[2]
        g[u] = append(g[u], Edge{v, w})
        g[v] = append(g[v], Edge{u, w})
    }

    dis := make([]int, n)
    for i := range dis {
        dis[i] = 1 << 31 - 1
    }
    dis[n-1] = 0

    h := &hp{Edge{n-1, 0}}
    for h.Len() > 0 {
        cur := heap.Pop(h).(Edge)
        u, w := cur.node, cur.weight
        if w > dis[u] {
            continue
        }

        for _, e := range g[u] {
            v, vw := e.node, e.weight
            if dis[u] + vw < dis[v] {
                dis[v] = dis[u] + vw
                heap.Push(h, Edge{v, dis[v]})
            }
        }
    }

    f := make([]int, n)
    f[n-1] = 1

    node := make([]int, n)
    for i := range node {
        node[i] = i
    }

    sort.Slice(node, func(i, j int) bool {
        return dis[node[i]] < dis[node[j]]
    })

    for _, u := range node {
        for _, e := range g[u] {
            v := e.node
            if dis[v] < dis[u] {
                f[u] = (f[v] + f[u]) % MOD
            }
        }
    }

    return f[0]
}