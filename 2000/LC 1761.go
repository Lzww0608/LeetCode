func minTrioDegree(n int, edges [][]int) (ans int) {
    g := make([][]int, n)
    m := make([]map[int]bool, n)
    deg := make([]int, n)
    for i := 0; i < n; i++ {
        m[i] = make(map[int]bool)
        m[i][i] = true
    }
    for _, e := range edges {
        u, v := e[0] - 1, e[1] - 1
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
        m[u][v] = true
        m[v][u] = true
        deg[v]++
        deg[u]++
    }
    f := false
    ans = math.MaxInt32
    for i := 0; i < n; i++ {
        for _, u := range g[i] {
            if u <= i {
                continue
            }
            for _, v := range g[u] {
                if v <= i || v <= u || !m[i][v] {
                    continue
                }
                f = true
                t := deg[i] + deg[u] + deg[v] - 6

                ans = min(ans, t)
            }
        }
    }

    if !f {
        return -1
    }

    return ans
}   