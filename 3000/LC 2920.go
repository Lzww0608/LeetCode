func maximumPoints(edges [][]int, coins []int, k int) int {
    n := len(edges) + 1
    g := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    memo := make([][]int, n)
    for i := 0; i < n; i++ {
        memo[i] = make([]int, 15)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int, int) int
    dfs = func(i, j, fa int) int {
        p := &memo[i][j]
        if *p != -1 {
            return *p 
        }
        a := (coins[i] >> j) - k
        b := coins[i] >> (j + 1)
        for _, u := range g[i] {
            if u == fa {
                continue
            }
            a += dfs(u, j, i)
            if j + 1 < 14 {
                b += dfs(u, j + 1, i)
            }
        }
        *p = max(a, b)
        return *p
    }

    return dfs(0, 0, -1)
}