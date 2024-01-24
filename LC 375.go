const N int = 205
var dp [N][N]int 
func getMoneyAmount(n int) int {
    for d := 2; d <= n; d++ {
        for l := 1; l + d - 1 <= n; l++ {
            r := l + d - 1
            dp[l][r] = 0x3f3f3f3f
            for x := l; x <= r; x++ {
                cur := max(dp[l][x-1], dp[x+1][r]) + x
                dp[l][r] = min(dp[l][r], cur)
            }
        }
    }
    return dp[1][n]
}
