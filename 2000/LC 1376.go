func numOfMinutes(n int, headID int, manager []int, informTime []int) (ans int) {
    g := make([][]int, n)
    for i, x := range manager {
        if x == -1 {
            continue
        }
        g[x] = append(g[x], i)
    }
    var dfs func(int, int)
    dfs = func(u, time int) {
        ans = max(ans, time)
        if len(g[u]) == 0 {
            return
        }
        for _, v := range g[u] {
            dfs(v, time + informTime[u])
        }
    }
    dfs(headID, 0)

    return
}