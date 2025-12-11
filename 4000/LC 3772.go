func maxSubgraphScore(n int, edges [][]int, good []int) []int {
    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    ans := make([]int, n)
    f := make([]int, n)
    var dfs1 func(int, int) int 
    dfs1 = func(u, fa int) int {
        d := 1
        if good[u] == 0 {
            d = -1
        }

        for _, v := range g[u] {
            if v == fa {
                continue
            }
            cur := dfs1(v, u)
            d += max(cur, 0)
        }

        f[u] = d
        return max(d, 0)
    }
    dfs1(0, -1)
    ans[0] = f[0]

    var dfs2 func(int, int, int)
    dfs2 = func(u, fa, d int) {
        ans[u] = f[u] + max(0, d)
        for _, v := range g[u] {
            if v == fa {
                continue
            }
            
            cur := ans[u]
            if f[v] > 0 {
                cur -= f[v]
            }
            dfs2(v, u, max(0, cur))
        }
    }
    dfs2(0, -1, 0)

    return ans
}