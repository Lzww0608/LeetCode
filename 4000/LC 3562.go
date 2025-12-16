func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
    e := make([][]int, n + 1)
    for _, v := range hierarchy {
        e[v[0]] = append(e[v[0]], v[1])
    }

    f := make([][][2]int, n + 1)
    for i := range f {
        f[i] = make([][2]int, budget + 1)
    }

    var dfs func(int) 
    dfs = func(u int) {
        g := make([][2]int, budget + 1)
        for i := range g {
            g[i] = [2]int{math.MinInt32, math.MinInt32}
            f[u][i] = [2]int{math.MinInt32, math.MinInt32}
        }
        g[0] = [2]int{0, 0}

        for _, v := range e[u] {
            dfs(v)

            for j := budget; j >= 0; j-- {
                for k := 0; k <= j; k++ {
                    g[j][0] = max(g[j][0], g[j - k][0] + f[v][k][0])
                    g[j][1] = max(g[j][1], g[j - k][1] + f[v][k][1])
                }
            }
        }

        for i := range budget + 1 {
            f[u][i][0] = g[i][0]
            f[u][i][1] = g[i][0]

            c, w := present[u - 1], future[u - 1] 
            if i >= c {
                f[u][i][0] = max(f[u][i][0], g[i - c][1] + w - c)
            }
            c /= 2 
            if i >= c {
                f[u][i][1] = max(f[u][i][1], g[i - c][1] + w - c)
            }
        }
    }
    dfs(1)

    ans := math.MinInt32
    for i := range budget + 1 {
        ans = max(ans, f[1][i][0])
    }

    return ans
}