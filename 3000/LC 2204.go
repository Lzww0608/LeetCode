func distanceToCycle(n int, edges [][]int) []int {
    ans := make([]int, n)
    g := make([][]int, n)
    deg := make([]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
        deg[u]++
        deg[v]++
    }

    q := []int{}
    for i, x := range deg {
        if x == 1 {
            q = append(q, i)
        }
    }

    s := make(map[int]bool)
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        s[u] = true
        for _, v := range g[u] {
            if deg[v]--; deg[v] == 1 {
                q = append(q, v)
            }
        }
    }

    var dfs func(int, int, int)
    dfs = func(u, fa, d int) {
        ans[u] = d
        for _, v := range g[u] {
            if v != fa && s[v] {
                dfs(v, u, d + 1)
            }
        }
    }
    for i := 0; i < n; i++ {
        if !s[i] {
            dfs(i, -1, 0)
        }
    }
    
    return ans
}