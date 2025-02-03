const MOD int = 1_000_000_007
func countOfArrays(n int, m int, k int) int {
    even, odd := m / 2, (m + 1) / 2
    f := make([][2]int, k + 1)
    f[0][0] = odd // odd
    f[0][1] = even // even

    for i := 1; i < n; i++ {         
        for j := k; j > 0; j-- {
            f[j][0], f[j][1] = (f[j][0] + f[j][1]) * odd % MOD, (f[j][0] + f[j-1][1]) * even % MOD
        }
        f[0][0], f[0][1] = (f[0][0] + f[0][1]) * odd % MOD, f[0][0] * even % MOD
    }

    return (f[k][0] + f[k][1]) % MOD
}