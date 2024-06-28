const MOD int = 1_000_000_007
func kInversePairs(n int, k int) int {
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, k + 1)
    }
    f[1][0] = 1

    for i := 2; i <= n; i++ {
        f[i][0] = 1
        for j := 1; j <= k; j++ {
            f[i][j] = (f[i-1][j] + f[i][j-1]) % MOD
            if j >= i {
                f[i][j] = (f[i][j] - f[i-1][j-i] + MOD) % MOD
            }
        } 
    }

    return f[n][k]
}