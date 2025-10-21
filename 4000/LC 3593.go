func minIncrease(n int, edges [][]int, cost []int) (ans int) {
    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }
    g[0] = append(g[0], -1)


    var dfs func(int, int) int
    dfs = func(u, fa int) int {
        if len(g[u]) == 1 {
            return cost[u]
        }
        cnt, mx := 0, 0
        for _, v := range g[u] {
            if v == fa {
                continue
            }
            x := dfs(v, u)
            if x > mx {
                cnt, mx = 1, x 
            } else if x == mx {
                cnt++
            }
        } 

        ans += len(g[u]) - 1 - cnt 
        return cost[u] + mx
    }
    dfs(0, -1)

    return 
}