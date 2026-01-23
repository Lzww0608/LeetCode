func minimumFlips(n int, edges [][]int, start string, target string) (ans []int) {
    type pair struct {
        e, i int
    }
    g := make([][]pair, n)
    for i, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], pair{v, i})
        g[v] = append(g[v], pair{u, i})
    }

    var dfs func(pair, int) int 
    dfs = func(u pair, fa int) int {
        x := 0
        for _, v := range g[u.e] {
            if v.e == fa {
                continue
            }
            x ^= dfs(v, u.e)
        }
        if (start[u.e] == target[u.e]) != (x == 0) {
            if u.i == -1 {
                ans = []int{-1}
                return -1
            }
            ans = append(ans, u.i)
            return 1
        }

        return 0
    }
    dfs(pair{0, -1}, -1)
    sort.Ints(ans)

    return
}