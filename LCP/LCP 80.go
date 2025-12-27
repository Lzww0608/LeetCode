func evolutionaryRecord(parents []int) string {
    n := len(parents)
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        j := parents[i]
        g[j] = append(g[j], i)
    }    

    var dfs func(int) string 
    dfs = func(u int) string {
        m := len(g[u])
        ans := make([]string, m)
        for i, v := range g[u] {
            ans[i] = dfs(v)
        }
        sort.Strings(ans)
        return "0" + strings.Join(ans, "") + "1"
    }

    return strings.TrimRight(dfs(0)[1:], "1")
}