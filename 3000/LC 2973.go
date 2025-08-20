func placedCoins(edges [][]int, cost []int) []int64 {
    n := len(edges) + 1
    ans := make([]int64, n)
    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    var dfs func(int, int) []int
    dfs = func(u, fa int) (a []int) {
        for _, v := range g[u] {
            if v != fa {
                a = append(a, dfs(v, u)...)
            }
        }
        a = append(a, cost[u])
        sort.Ints(a)
        if len(a) < 3 {
            ans[u] = 1
            return 
        }
        m := len(a)
        mx := max(a[0] * a[1] * a[m - 1], a[m - 1] * a[m - 2] * a[m - 3])
        if mx > 0 {
            ans[u] = int64(mx)
        }

        if m > 5 {
            a = append(a[:2], a[m - 3:]...)
        }
        return
    }
    dfs(0, -1)

    return ans
}