func criticalConnections(n int, connections [][]int) (ans [][]int) {
    g := make([][]int, n)
    low := make([]int, n)
    dfn := make([]int, n)
    k := 0
    for _, e := range connections {
        u, v := e[0], e[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    var tarjan func(u, fa int)
    tarjan = func(u, fa int) {
        k++
        dfn[u], low[u] = k, k

        for _, v := range g[u] {
            if v == fa {
                continue
            }

            if dfn[v] == 0 {
                tarjan(v, u)
                low[u] = min(low[u], low[v])
                if dfn[u] < low[v] {
                    ans = append(ans, []int{u, v})
                }
            } else {
                low[u] = min(low[u], dfn[v])
            }
        }
    }

    tarjan(0, -1)

    return
}