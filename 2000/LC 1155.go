const MOD int = 1_000_000_007
func numRollsToTarget(n int, k int, target int) int {
    if target < n || target > n * k {
        return 0
    }
    f := make([]int, target + 1)
    f[0] = 1
    for i := 1; i <= n; i++ {
        g := make([]int, target + 1)
        sum := 0
        mx := min(i * k, target)
        for j := 1; j <= mx; j++ {
            sum = (sum + f[j-1]) % MOD
            if j > k {
                sum = (sum - f[j-k-1] + MOD) % MOD
            }
            g[j] = sum
        }
        f = g
    }

    return f[target]
}