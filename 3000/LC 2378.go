func maxScore(edges [][]int) int64 {
    type node struct {
        u, val int
    }
    n := len(edges) + 1
    g := make([][]node, n)
    for u, edge := range edges {
        if u == 0 {
            continue
        }
        v, x := edge[0], edge[1]
        g[v] = append(g[v], node{u, x})
    }

    var dfs func(int) (int, int)
    dfs = func(u int) (a, b int) {
        if len(g[u]) == 0 {
            return 
        }

        for _, v := range g[u] {
            x, y := dfs(v.u)
            a += max(x, y)
            b = max(b, x + v.val - max(x, y))
        }
        b += a

        return
    }

    return int64(max(dfs(0)))
}