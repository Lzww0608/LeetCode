func sellingWood(m int, n int, prices [][]int) int64 {
    dp := [205][205]int64{}

    for _, p := range prices {
        a, b, c := p[0], p[1], p[2]
        dp[a][b] = int64(c)
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            for k := 1; k < i; k++ {
                dp[i][j] = max(dp[i][j], dp[i-k][j] + dp[k][j])
            }

            for k := 1; k < j; k++ {
                dp[i][j] = max(dp[i][j], dp[i][k] + dp[i][j-k])
            }
        }
    }

    return dp[m][n]
}
