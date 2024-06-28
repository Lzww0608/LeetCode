var MOD int = int(1e9 + 7)
func checkRecord(n int) int {
    f := make([][2][3]int, n + 1)
    f[0][0][0] = 1

    for i := 0; i < n; i++ {
        for j := 0; j < 2; j++ {
            for k := 0; k < 3; k++ {
                if j != 1 {
                    f[i + 1][j + 1][0] = (f[i + 1][j + 1][0] + f[i][j][k]) % MOD
                }
                if k != 2 {
                    f[i + 1][j][k + 1] = (f[i + 1][j][k + 1] + f[i][j][k]) % MOD
                }
                f[i + 1][j][0] = (f[i + 1][j][0] + f[i][j][k]) % MOD
            }
        }
    }

    res := 0
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            res = (res + f[n][i][j]) % MOD
        }
    }

    return res
}
