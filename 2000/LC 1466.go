func minReorder(n int, connections [][]int) (ans int) {
    type pair struct {
        u, b int 
    }

    g := make([][]pair, n)
    for _, connection := range connections {
        u, v := connection[0], connection[1]
        g[u] = append(g[u], pair{v, 0})
        g[v] = append(g[v], pair{u, 1})
    }

    var dfs func(int, int)
    dfs = func(u, fa int) {
        for _, v := range g[u] {
            if v.u == fa {
                continue
            }
            if v.b == 0 {
                ans++
            }
            dfs(v.u, u)
        }
    }
    dfs(0, -1)

    return
}