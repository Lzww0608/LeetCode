func makeConnected(n int, connections [][]int) (ans int) {
    m := len(connections)
    if m < n - 1 {
        return -1
    }

    g := make([][]int, n)
    for _, connection := range connections {
        a, b := connection[0], connection[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    vis := make([]bool, n)

    var dfs func(int, int)
    dfs = func(u, fa int) {
        if vis[u] {
            return 
        }
        vis[u] = true
        for _, v := range g[u] {
            if v != fa {
                dfs(v, u)
            }
        }
        return
    }

    for i := 0; i < n; i++ {
        if !vis[i] {
            ans++
            dfs(i, -1)
        }
    }

    return ans - 1
}