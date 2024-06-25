func minimumDeleteSum(s1 string, s2 string) int {
    n, m := len(s1), len(s2)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    }

    // Initialize base cases
    for i := n - 1; i >= 0; i-- {
        dp[i][m] = dp[i+1][m] + int(s1[i])
    }
    for j := m - 1; j >= 0; j-- {
        dp[n][j] = dp[n][j+1] + int(s2[j])
    }

    // Fill in dp table
    for i := n - 1; i >= 0; i-- {
        for j := m - 1; j >= 0; j-- {
            if s1[i] == s2[j] {
                dp[i][j] = dp[i+1][j+1]
            } else {
                dp[i][j] = min(dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j]))
            }
        }
    }

    return dp[0][0]
}

