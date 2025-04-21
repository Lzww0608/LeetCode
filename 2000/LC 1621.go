const MOD int = 1_000_000_007
func numberOfSets(n int, k int) int {
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, k + 1)
        f[i][0] = 1
    }

    // f[i][j] = f[i-1][j] + f[i-1][j-1] + ... + f[1][j-1]
    // f[i-1][j] = f[i-2][j] + f[i-2][j-1] + ... + f[1][j-1]
    // f[i][j] - f[i-1][j] = f[i-1][j] - f[i-2][j] + f[i-1][j-1]
    for i := 2; i <= n; i++ {
        for j := 1; j <= min(i - 1, k); j++ {
            f[i][j] = f[i-1][j] * 2 + f[i-1][j-1] - f[i-2][j] + MOD
            f[i][j] %= MOD
        }
    }

    return f[n][k]
}