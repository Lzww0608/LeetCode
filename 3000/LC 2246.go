func longestPath(parent []int, s string) (ans int) {
    n := len(parent)
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        x := parent[i]
        g[x] = append(g[x], i)
    }

    var dfs func(int) int
    dfs = func(x int) (res int) {
        for _, y := range g[x] {
            t := dfs(y) + 1
            if s[x] != s[y] {
                ans = max(ans, res + t)
                res = max(res, t)
            }
        }
        return
    }
    dfs(0)
    return ans + 1
}