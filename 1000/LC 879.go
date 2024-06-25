var MOD int = int(1e9 + 7)
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, minProfit + 1)
        dp[i][0] = 1
    }
    for i := range group {
        for j := n; j >= group[i]; j-- {
            for k := minProfit; k >= 0; k-- {
                dp[j][k] += dp[j-group[i]][max(k-profit[i], 0)]
                dp[j][k] %= MOD
            }
        }
    }
    return dp[n][minProfit]
    
}
