func solveNQueens(n int) [][]string {
    ans, res := make([][]string,0), make([][]rune,0)
    for i := 0; i < n; i++ {
        s := []rune{}
        for j := 0; j < n; j++ {
            s = append(s, '.')
        }
        res = append(res, s)
    }
    col, dg, udg := [20]bool{}, [20]bool{}, [20]bool{}
    var dfs func(x int)
    dfs = func(x int) {
        if x == n {
            var s []string
            for _, t := range res {
                s = append(s, string(t))
            }
            ans = append(ans, s)
            return
        }
        for i := 0; i < n; i++ {
            if !col[i] && !dg[x+i] && !udg[n-x+i] {
                res[x][i] = 'Q'
                col[i] = true
                dg[x+i] = true
                udg[n-x+i] = true
                dfs(x+1)
                res[x][i] = '.'
                col[i] = false
                dg[x+i] = false
                udg[n-x+i] = false
            }
        }
    }
    dfs(0)
    return ans
}