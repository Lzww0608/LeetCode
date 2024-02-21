var MOD int = int(1e9 + 7)
func numTilings(n int) int {
    dp := make([][4]int, n+1)
    dp[0][3] = 1
    for i := 1; i <= n; i++ {
        dp[i][0] = dp[i-1][3]
        dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % MOD
        dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % MOD
        dp[i][3] = (dp[i-1][0] + dp[i-1][1]+ dp[i-1][2] + dp[i-1][3]) % MOD
    } 
    return dp[n][3]
}