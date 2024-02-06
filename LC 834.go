func sumOfDistancesInTree(n int, edges [][]int) []int {
    g := make([][]int, n)
    ans := make([]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }
    dp := make([]int, n)
    var dfs1 func(v, fa, depth int)
    dfs1 = func(v, fa, depth int) {
        ans[0] += depth
        dp[v] = 1
        for _, u := range g[v] {
            if (u == fa) {
                continue
            }
            dfs1(u, v, depth+1)
            dp[v] += dp[u]
        }
    }
    
    var dfs2 func(v, fa int) 
    dfs2 = func(v, fa int) {
        for _, u := range g[v] {
            if fa != u {
                ans[u] = ans[v] + n - 2 * dp[u]
                dfs2(u, v)
            }
        }
    }

    dfs1(0, -1, 0)
    dfs2(0, -1)
    return ans
}