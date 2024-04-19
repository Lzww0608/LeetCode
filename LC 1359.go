const MOD int = 1e9 + 7
func countOrders(n int) int {
    ans := 1
    for i := 1; i <= n; i++ {
        ans = ans * ((2 * i - 1) * 2 * i) / 2 % MOD
    }

    return ans
}