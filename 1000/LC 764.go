func orderOfLargestPlusSign(n int, mines [][]int) (ans int) {
    g := make([][]int, n + 2)
    l := make([][]int, n + 2)
    r := make([][]int, n + 2)
    u := make([][]int, n + 2)
    d := make([][]int, n + 2)
    for i := range g {
        g[i] = make([]int, n + 2)
        l[i] = make([]int, n + 2)
        r[i] = make([]int, n + 2)
        u[i] = make([]int, n + 2)
        d[i] = make([]int, n + 2)
        for j := range g {
            g[i][j] = 1
        }
    }
    for _, v := range mines {
        g[v[0]+1][v[1]+1] = 0 
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if g[i][j] == 1 {
                u[i][j] = u[i-1][j] + 1
                l[i][j] = l[i][j-1] + 1
            }
            if g[n-i+1][n-j+1] == 1 {
                d[n-i+1][n-j+1] = d[n-i+2][n-j+1] + 1
                r[n-i+1][n-j+1] = r[n-i+1][n-j+2] + 1
            }
        }
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            ans = max(ans, min(l[i][j], r[i][j], u[i][j], d[i][j]))
        }
    }

    return
}