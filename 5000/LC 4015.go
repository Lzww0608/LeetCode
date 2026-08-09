func weightedSum(parent []int, nums []int) int64 {
    n := len(parent)
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        g[parent[i]] = append(g[parent[i]], i)
    }

    var dfs func(int, int) int 
    dfs = func(u, fa int) int {
        d := 1
        for _, v := range g[u] {
            if v != fa {
                d = max(d, 1 + dfs(v, u))
            }
        }
        return d
    }

    h := dfs(0, -1)
    ans := 0
    var dfs1 func(int, int, int) 
    dfs1 = func(u, fa, d int) {
        ans += nums[u] * (h - d + 1)
        for _, v := range g[u] {
            if v != fa {
                dfs1(v, u, d + 1)
            }
        }
    }
    dfs1(0, -1, 1)

    return int64(ans)
}
