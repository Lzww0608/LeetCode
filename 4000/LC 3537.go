func specialGrid(n int) [][]int {
    m := 1 << n 
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, m)
    }

    k := 0
    var dfs func(l, r, u, d int)
    dfs = func(l, r, u, d int) {
        if d - u == 1 {
            ans[u][l] = k 
            k++
            return
        }

        m := (d - u) / 2
        dfs(l + m, r, u, u + m)
        dfs(l + m, r, u + m, d)
        dfs(l, l + m, u + m, d)
        dfs(l, l + m, u, u + m)
    }

    dfs(0, m, 0, m)
    return ans
}