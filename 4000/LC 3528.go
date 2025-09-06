const MOD int = 1_000_000_007

type pair struct {
    u, w int
}

func baseUnitConversions(conversions [][]int) []int {
    n := len(conversions) + 1
    g := make([][]pair, n)
    for _, conversion := range conversions {
        u, v, w := conversion[0], conversion[1], conversion[2]
        g[u] = append(g[u], pair{v, w})
        g[v] = append(g[v], pair{u, w})
    }

    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }

    var dfs func(u, fa, w int) 
    dfs = func(u, fa, w int) {
        if ans[u] != -1 {
            return
        }
        ans[u] = w % MOD
        for _, v := range g[u] {
            if v.u == fa {
                continue
            }
            dfs(v.u, u, w * v.w % MOD)
        }
    }
    dfs(0, -1, 1)

    return ans
}