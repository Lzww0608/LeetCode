const MOD int = 1_000_000_007
func assignEdgeWeights(edges [][]int) int {
    n := len(edges) + 1
    g := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0] - 1, edge[1] - 1
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    var dfs func(int, int) int 
    dfs = func(u, fa int) (res int) {
        for _, v := range g[u] {
            if v != fa {
                res = max(res, 1 + dfs(v, u))
            }
        }

        return
    }

    k := dfs(0, -1)

    return quickPow(2, k - 1)
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        r >>= 1
        a = a * a % MOD
    }

    return res
}