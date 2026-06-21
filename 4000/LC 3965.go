func finishTime(n int, edges [][]int, baseTime []int) int64 {
    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    if n == 1 {
        return int64(baseTime[0])
    }

    var dfs func(int, int) int 
    dfs = func(u, fa int) int {
        if u != 0 && len(g[u]) == 1 {
            return baseTime[u]
        }

        mx, mn := math.MinInt, math.MaxInt
        for _, v := range g[u] {
            if v == fa {
                continue
            }
            cur := dfs(v, u)
            mx = max(mx, cur)
            mn = min(mn, cur)
        } 
        return mx * 2 - mn + baseTime[u]
    }

    return int64(dfs(0, -1))
}