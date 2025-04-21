const MOD int = 1_000_000_007
func numMusicPlaylists(n int, goal int, k int) int {
    f := make([][]int, goal + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }

    f[0][0] = 1
    for i := 1; i <= goal; i++ {
        for j := 1; j <= n; j++ {
            f[i][j] = (f[i-1][j] * max(0, j - k) + f[i-1][j-1] * (n - j + 1)) % MOD
        }
    }

    return f[goal][n]
}