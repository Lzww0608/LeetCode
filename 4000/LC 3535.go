const MOD int = 1_000_000_007

func inverse(a, b int) int {
    return a * quickPow(b, MOD - 2) % MOD
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

type pair struct {
    u, w int
}

func queryConversions(conversions [][]int, queries [][]int) []int {
    n := len(conversions) + 1
    g := make([][]pair, n)
    for _, conversion := range conversions {
        u, v, w := conversion[0], conversion[1], conversion[2]
        g[u] = append(g[u], pair{v, w})
        g[v] = append(g[v], pair{u, w})
    }

    a := make([]int, n)
    for i := range a {
        a[i] = -1
    }

    var dfs func(u, fa, w int) 
    dfs = func(u, fa, w int) {
        if a[u] != -1 {
            return
        }
        a[u] = w % MOD
        for _, v := range g[u] {
            if v.u == fa {
                continue
            }
            dfs(v.u, u, w * v.w % MOD)
        }
    }
    dfs(0, -1, 1)

    ans := make([]int, len(queries))
    for i, q := range queries {
        u, v := q[0], q[1]
        ans[i] = inverse(a[v], a[u])
    }

    return ans
}