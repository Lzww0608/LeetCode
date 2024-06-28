func soupServings(n int) float64 {
    n = (n + 24) / 25
    if n >= 800 {
        return 1
    }

    dp := make([][]float64, n + 1)
    for i := range dp {
        dp[i] = make([]float64, n + 1)
    }

    var dfs func(int, int) float64
    dfs = func(i, j int) float64 {

        if i <= 0 && j <= 0 {
            return 0.5
        } else if i <= 0 {
            return 1
        } else if j <= 0 {
            return 0
        }
        if dp[i][j] == 0 {
            dp[i][j] = (dfs(i - 4, j) + dfs(i - 3, j - 1) + dfs(i - 2, j - 2) + dfs(i - 1, j - 3)) / 4
        }
        return dp[i][j]
    }
    return dfs(n, n)

}
