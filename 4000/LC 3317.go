const MOD int = 1_000_000_007
func numberOfWays(n int, x int, y int) (ans int) {
    s := make([][]int, n + 1)
    for i := range s {
        s[i] = make([]int, n + 1)
    }
    s[0][0] = 1

    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            s[i][j] = (s[i - 1][j - 1] + j * s[i - 1][j]) % MOD
        }
    }

    perm, score := 1, 1
    for i := 1; i <= n; i++ {
        perm = perm * (x - i + 1) % MOD
        score = score * y % MOD
        ans = (ans + perm * score % MOD * s[n][i] % MOD) % MOD
    }

    return 
}