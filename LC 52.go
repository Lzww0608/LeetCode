func totalNQueens(n int) int {
    ans := 0
    col, dg, udg := [20]bool{}, [20]bool{}, [20]bool{}
    var dfs func(x int)
    dfs = func(x int) {
        if x == n {
            ans++
            return
        }
        for i := 0; i < n; i++ {
            if !col[i] && !dg[x+i] && !udg[n-i+x] {
                col[i] = true
                dg[x+i] = true
                udg[n-i+x] = true
                dfs(x+1)
                col[i] = false
                dg[x+i] = false
                udg[n-i+x] = false
            }
        }
    }
    dfs(0)
    return ans
}