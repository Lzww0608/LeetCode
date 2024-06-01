func minimumOperations(leaves string) int {
    n := len(leaves)
    
    dp := make([][3]int, n)
    dp[0][0] = boolToInt(leaves[0] == 'y')
    dp[0][1] = int(1e9)
    dp[0][2] = int(1e9)
    dp[1][2] = int(1e9)
    
    for i := 1; i < n; i++ {
        dp[i][0] = dp[i-1][0] + boolToInt(leaves[i] == 'y')
        dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + boolToInt(leaves[i] == 'r')
        if i > 1 {
            dp[i][2] = min(dp[i-1][1], dp[i-1][2]) + boolToInt(leaves[i] == 'y')
        }
    }

    return dp[n-1][2]
}

func boolToInt(b bool) int {
    if b {
        return 1
    }
    return 0
}