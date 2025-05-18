// 0 1 2 9 44 265
const MOD int = 1_000_000_007
func findDerangement(n int) int {
    f := make([]int, n + 2)
    f[1], f[2] = 0, 1
    for i := 3; i <= n; i++ {
        f[i] = (f[i-2] + f[i-1]) * (i - 1) % MOD
    }
    return f[n]
}