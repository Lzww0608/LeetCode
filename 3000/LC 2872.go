func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) (ans int) {
    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    var dfs func(int, int) int 
    dfs = func(u, fa int) (sum int) {
        sum += values[u]
        for _, v := range g[u] {
            if v != fa {
                sum += dfs(v, u)
            }
        }
        if sum % k == 0 {
            ans++
        }

        return sum
    }
    dfs(0, -1)

    return 
}