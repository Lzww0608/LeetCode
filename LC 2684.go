func maxMoves(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    for j := n - 2; j >= 0; j-- {
        for i := 0; i < m; i++ {
            if i - 1 >= 0 && grid[i][j] < grid[i-1][j+1] {
                dp[i][j] = max(dp[i][j], dp[i-1][j+1] + 1)
            } 
            if grid[i][j] < grid[i][j+1] {
                dp[i][j] = max(dp[i][j], dp[i][j+1] + 1)
            }
            if i + 1 < m && grid[i][j] < grid[i+1][j+1] {
                dp[i][j] = max(dp[i][j], dp[i+1][j+1] + 1)
            } 
        }
    }

    for i := 0; i < m; i++ {
        ans = max(ans, dp[i][0])
    }
    
    return
}