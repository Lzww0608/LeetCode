func solveNQueens(n int) (ans [][]string) {
    g := make([][]rune, n)
    s := make([]rune, n)
    for i := range s {
        s[i] = '.'
    }
    
    for i := range g {
        g[i] = append([]rune(nil), s...)
    }

    var col, dg, udg [20]bool
    var dfs func(int)
    dfs = func(x int) {
        if x == n {
            var tmp []string
            for i := range g {
                tmp = append(tmp, string(g[i]))
            }
            ans = append(ans, tmp)
            return
        }

        for i := 0; i < n; i++ {
            if !col[i] && !dg[x+i] && !udg[i-x+n] {
                col[i] = true
                dg[x+i] = true
                udg[i-x+n] = true
                g[x][i] = 'Q'
                dfs(x+1)
                g[x][i] = '.'
                col[i] = false
                dg[x+i] = false
                udg[i-x+n] = false
            }
        }
    }
    dfs(0)

    return
}