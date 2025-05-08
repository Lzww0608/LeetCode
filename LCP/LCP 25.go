const MOD int = 1_000_000_007
func keyboard(k int, n int) int {
    f := make([][27]int, n + 1)
    for i := 0; i <= 26; i++ {
        f[0][i] = 1
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= 26; j++ {
            for p := 0; p <= min(i, k); p++ {
                f[i][j] = (f[i][j] + f[i-p][j-1] * C(i, p)) % MOD
            }
        }
    }

    return f[n][26]
}

func C(m, n int) int {
    ans, k := 1, 1
    for k <= n {
        ans = (ans * (m - k + 1)) / k % MOD
        k++
    }

    return ans
}