func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
    n := len(edges) + 1
    g := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    g[0] = append(g[0], -1)
    var dfs func(int, int) int 
    dfs = func(u, fa int) int {
        if len(g[u]) == 1 {
            return values[u]
        }

        cur := 0
        for _, v := range g[u] {
            if v != fa {
                cur += dfs(v, u)
            }
        }

        return min(values[u], cur)
    }

    sum := 0
    for _, x := range values {
        sum += x
    }

    return int64(sum - dfs(0, -1))
}