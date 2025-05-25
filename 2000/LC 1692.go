const MOD int = 1_000_000_007

func waysToDistribute(n int, k int) int {
    f := make([][]int, k + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
        f[i][0] = 1
    }

    for i := 1; i <= k; i++ {
        f[i][i] = 1
        for j := i + 1; j <= n; j++ {
            f[i][j] = (f[i-1][j-1] + f[i][j-1] * i) % MOD
        }
    }

    return f[k][n]
}