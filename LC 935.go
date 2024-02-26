var MOD int = int(1e9 + 7)
func knightDialer(n int) int {
    dp := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
    for i := 1; i < n; i++ {
        tmp := make([]int, 10)
        tmp[0] = (dp[6] + dp[4]) % MOD
        tmp[1] = (dp[8] + dp[6]) % MOD
        tmp[2] = (dp[7] + dp[9]) % MOD
        tmp[3] = (dp[8] + dp[4]) % MOD
        tmp[4] = (dp[3] + dp[9] + dp[0]) % MOD
        tmp[6] = (dp[1] + dp[7] + dp[0]) % MOD
        tmp[7] = (dp[6] + dp[2]) % MOD
        tmp[8] = (dp[1] + dp[3]) % MOD
        tmp[9] = (dp[2] + dp[4]) % MOD
        copy(dp, tmp)
    }
    ans := 0
    for _, x := range dp {
        ans = (ans + x) % MOD
    }
    return ans
}